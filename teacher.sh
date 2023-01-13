$var_E = head -n 179 streets/Buckingham_Place | tail -n 1
echo $var_E
cat interviews/interview-"$var_E"
echo $MAIN_SUSPECT 