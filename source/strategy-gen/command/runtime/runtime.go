package runtime

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tsinghua-cel/strategy-gen/library"
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
	"net/http"
	"time"
)

func GetCommand() *cobra.Command {
	runtimeCmd := &cobra.Command{
		Use:   "runtime",
		Short: "Run a library strategy generator.",
		Run:   runCommand,
	}
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	setFlags(runtimeCmd)
	return runtimeCmd
}

func setFlags(cmd *cobra.Command) {

	cmd.Flags().IntVar(
		&params.maxValidatorIndex,
		maxValidatorIndexFlag,
		-1,
		"the max hack validator index",
	)

	cmd.Flags().StringVar(
		&params.attacker,
		attackerFlag,
		"127.0.0.1:12001",
		"the attacker service to update",
	)
}

func runCommand(cmd *cobra.Command, _ []string) {
	var latestEpoch int64
	ticker := time.NewTicker(time.Second * 3)
	slotTool := utils.SlotTool{SlotsPerEpoch: 32}
	for {
		select {
		case <-ticker.C:
			slot, err := utils.GetSlot(params.attacker)
			if err != nil {
				log.WithField("error", err).Error("failed to get slot")
				continue
			}
			epoch := slotTool.SlotToEpoch(int64(slot))
			if epoch == latestEpoch {
				continue
			}
			if int64(slot) < slotTool.EpochEnd(epoch) {
				continue
			}
			latestEpoch = epoch
			// get next epoch duties
			duties, err := utils.GetEpochDuties(params.attacker, epoch+1)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
					"epoch": epoch + 1,
				}).Error("failed to get duties")
				latestEpoch = epoch - 1
				continue
			}
			if hackDuties, happen := library.CheckDuties(params.maxValidatorIndex, duties); happen {
				strategy := types.Strategy{}
				strategy.Validators = library.ValidatorStrategy(hackDuties)
				strategy.Slots = library.GenSlotStrategy(hackDuties)
				if err = updateStrategy(params.attacker, strategy); err != nil {
					log.WithField("error", err).Error("failed to update strategy")
				} else {
					log.WithFields(log.Fields{
						"epoch":    epoch + 1,
						"strategy": strategy,
					}).Info("update strategy successfully")
				}
			}
		}
	}
}

func updateStrategy(url string, strategy types.Strategy) error {
	d, err := json.Marshal(strategy)
	if err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("http://%s/v1/update-strategy", url), "application/json", bytes.NewReader(d))
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("failed to generate strategy: %s", res.Status)
	}
	return nil
}
