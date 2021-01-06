{{ $channel := 787863951446245397 }}
{{$args := parseArgs 1  ""
  (carg "string" "value")}}
{{ $msg := ($args.Get 0) }}
 
{{sendMessageNoEscape $channel $msg}}
{{deleteTrigger 5}}