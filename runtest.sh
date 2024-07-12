#!/bin/bash
casetype=${1:-"1"}
caseduration=${2:-"600"}

basedir=$(pwd)
casedir="${basedir}/case"

testcase1() {
	subdir="blockcost"
	targetdir="${casedir}/${subdir}"
	resultdir="${basedir}/results/${subdir}"
	# if resultdir exist, delete it.
	if [ -d $resultdir ]; then
		rm -rf $resultdir
	fi
	mkdir -p $resultdir

	echo "Running testcase $subdir"
	echo "first test with normal version"
	docker compose -f $targetdir/docker-compose-normal.yml up -d 
	echo "wait $caseduration seconds" && sleep $caseduration
	docker compose -f $targetdir/docker-compose-normal.yml down
	echo "result collect"
	sudo mv data $resultdir/data-normal
	cd $resultdir/data-normal
	find . -name GetBeaconBlock.csv | xargs cat > /tmp/_b.csv
	sort -t "," -k 1n,1 /tmp/_b.csv > cost.csv
	cd $basedir

	echo "second test with reorg-fix version"
	docker compose -f $targetdir/docker-compose-reorg.yml up -d
	echo "wait $caseduration seconds" && sleep $caseduration
	docker compose -f $targetdir/docker-compose-reorg.yml down
	echo "result collect"
	sudo mv data $resultdir/data-reorg
	cd $resultdir/data-reorg
	find . -name GetBeaconBlock.csv | xargs cat > /tmp/_b.csv
	sort -t "," -k 1n,1 /tmp/_b.csv > cost.csv
	cd $basedir
	echo "test done and result in $resultdir"
}

testcase2() {
	subdir="tpstest"
	targetdir="${casedir}/${subdir}"
	resultdir="${basedir}/results/${subdir}"
	# if resultdir exist, delete it.
	if [ -d $resultdir ]; then
		rm -rf $resultdir
	fi
	mkdir -p $resultdir

	echo "Running testcase $subdir"
	echo "first test with normal version"
	docker compose -f $targetdir/docker-compose-normal.yml up -d 
	echo "wait 120 seconds" && sleep 120
	echo "begin txpress send"
	./txpress.sh
	docker compose -f $targetdir/docker-compose-normal.yml down
	sudo mv data $resultdir/data-normal

	echo "second test with reorg-fix version"
	docker compose -f $targetdir/docker-compose-reorg.yml up -d
	echo "wait 120 seconds" && sleep 120
	echo "begin txpress send"
	./txpress.sh
	docker compose -f $targetdir/docker-compose-reorg.yml down
	sudo mv data $resultdir/data-reorg
	echo "test done and result in $resultdir"
}


testcase3() {
	subdir="attack-tpstest"
	targetdir="${casedir}/${subdir}"
	resultdir="${basedir}/results/${subdir}"
	# if resultdir exist, delete it.
	if [ -d $resultdir ]; then
		rm -rf $resultdir
	fi
	mkdir -p $resultdir

	epochsToWait=7

	echo "Running testcase $subdir"
	echo "first test with normal version"
	docker compose -f $targetdir/docker-compose-normal.yml up -d 
	echo "wait $epochsToWait epochs" && sleep $(($epochsToWait * 12 * 32))
	echo "begin txpress send"
	./txpress.sh
	docker compose -f $targetdir/docker-compose-normal.yml down
	sudo mv data $resultdir/data-normal

	echo "second test with reorg-fix version"
	docker compose -f $targetdir/docker-compose-reorg.yml up -d
	echo "wait $epochsToWait epochs" && sleep $(($epochsToWait * 12 * 32))
	echo "begin txpress send"
	./txpress.sh
	docker compose -f $targetdir/docker-compose-reorg.yml down
	sudo mv data $resultdir/data-reorg
	echo "test done and result in $resultdir"
}

testcase4() {
	subdir="attack-reorg"
	targetdir="${casedir}/${subdir}"
	resultdir="${basedir}/results/${subdir}"
	# if resultdir exist, delete it.
	if [ -d $resultdir ]; then
		rm -rf $resultdir
	fi
	mkdir -p $resultdir

	epochsToWait=20

	echo "Running testcase $subdir"
	echo "first test with normal version"
	docker compose -f $targetdir/docker-compose-normal.yml up -d 
	echo "wait $epochsToWait epochs" && sleep $(($epochsToWait * 12 * 32))
	docker compose -f $targetdir/docker-compose-normal.yml down
	sudo mv data $resultdir/data-normal

	echo "second test with reorg-fix version"
	docker compose -f $targetdir/docker-compose-reorg.yml up -d
	echo "wait $epochsToWait epochs" && sleep $(($epochsToWait * 12 * 32))
	docker compose -f $targetdir/docker-compose-reorg.yml down
	sudo mv data $resultdir/data-reorg

	echo "test done and result in $resultdir"
}

switchcase() {
	case $casetype in
		1)
			testcase1
			;;
		2)
			testcase2
			;;
		3)
			testcase3
			;;
		4)
			testcase4
			;;
		*)
			echo "Invalid case type"
			;;
	esac
}
