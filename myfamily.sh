curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --argjson HeroID "$HERO_ID" '.[] | select( .id ==$HeroID) | .connections.relatives' | sed 's/"//g'
