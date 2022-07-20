
ADDR="localhost:9001/test3"
DATA='[{"a":"99","b":"2","key":"x"},{"a":"10","b":"1","c":"2","key":"y"}]' 
#DATA=$1

echo
echo $ADDR
echo

curl -X -POST -i -g \
-H 'Content-Type: application/json' \
-d $DATA \
$ADDR

echo
echo
