KEY=$1
VALUE=$2

DATA='{"key":"'$1'","value":"'$2'"}'

ADDR="localhost:9001/test1"

echo
echo "$ADDR"
echo $DATA
echo

curl -X -POST -i \
-H 'Content-Type: application/json' \
-d $DATA \
$ADDR

echo
echo
