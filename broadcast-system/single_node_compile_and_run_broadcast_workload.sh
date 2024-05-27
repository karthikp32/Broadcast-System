go install .
cd maelstrom/ && ./maelstrom test -w broadcast --bin $GOBIN/broadcast --node-count 1 --time-limit 20 --rate 10
