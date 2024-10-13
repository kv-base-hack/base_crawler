#!/usr/bin/env bash



abigen -abi ./swap_topic_1.abi -pkg contracts -type SwapTopic1 -out ./swap_topic_1.go 
abigen -abi ./swap_topic_2.abi -pkg contracts -type SwapTopic2 -out ./swap_topic_2.go 
abigen -abi ./swap_topic_6.abi -pkg contracts -type SwapTopic6 -out ./swap_topic_6.go

abigen -abi ./transfer_1.abi -pkg contracts -type Transfer1 -out ./transfer_1.go 



