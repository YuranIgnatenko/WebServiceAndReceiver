ADDR="localhost:9001/test2"
DATA='{"s":"'$1'","key":"'$2'"}'

echo
echo $ADDR
echo $DATA
echo

curl -X -POST -i \
-H 'Content-Type: application/json' \
-d $DATA \
$ADDR

echo
echo
