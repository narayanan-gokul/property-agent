notify-send -u low "Property checker" "Starting process"

temp="/home/gokul/automation/property-agent/temp"
./exec 2> /home/gokul/automation/property-agent/log/log.txt

ls_result=$(ls /home/gokul/automation/property-agent/temp | grep properties | tail -n 2)
IFS=$'\n' read -r -d '' -a files <<< "$ls_result"

if [ ${#files[@]} -eq 1 ]; then
	notify-send -u critical "Property checker" "Updates available\nOpen ~/my_obsidian_vault/projects/sydney/properties.md for details"
else
	notify-send "Property checker" "Comparing results"
	diff --color=auto $temp/"${files[0]}" $temp/"${files[1]}"
	diff_result=$?
	if [ $diff_result == 1 ]
	then
		notify-send -u critical "Property checker" "Updates available\nOpen ~/my_obsidian_vault/projects/sydney/properties.md for details"
	else
		notify-send -u low "Property checker" "No changes"
	fi
	rm $temp/"${files[0]}"
fi
cp $temp/"${files[1]}" /home/gokul/my_obsidian_vault/projects/sydney/properties.md
