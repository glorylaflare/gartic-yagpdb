{{/*
	# ESSE COMANDO NÃO FAZ MAIS PARTE DO SERVIDOR DO GARTIC, PORÉM FUNCIONA PERFEITAMENTE

	Comando que envia mensagens a cada intervalo de tempo para avisar sobre beber água e se hidratar 
 	
	Trigger type: Minute interval
	Interval: 30
	Channel: Qualquer um da sua escolha, pois o bot faz um sorteio para enviar a mensagens nos canais selecionados no $listCh
*/}}

{{ $listCh := cslice 559111535490760716 618578571786715193 487342947025813505 517551594103898116 517551615696306186 559111496634728448 742247486722146356 747279523913007125 758809541964988446 742246194821988364 752319253658730597 763920918623158322 746963919188263012 754091659645550652 750751981588447322 744461467277066307 742248173958856734 745046029530824846 758809592225333268 746963208103002207 742434451702808577 750433777473290288 771388309984903208 781992106645520434 763078891597201438 763078904659050536 763078833465196614 488080553359179797 528696634880819210 528803754728751104 757395492782080030 529221438452924426 747646659952902185 663136041405906989 750845599233867864 762683813226610706 456915295940902953 788996947423920128 788997005523288135 788997044650901518 768988994586279936 }}
{{ $listmsg := (cslice "**Que tal uma pausa para beber água?**" "**Já bebeu água hoje? Vai lá, a gente te espera.**" "**Da uma pausinha e vai se hidratar, faz bem para saúde.**" "**Um copo d'água cairia bem agora em? Que tal se hidratar um pouco?**" "**Para um pouco e vai beber água, você precisa se hidratar.**" "**Ta na hora de beber água em, só um pouquinho, vai lá.**" "**Beber água faz seu corpo funcionar melhor, vai lá, bebe um pouco.**" "**Melhor uma pedra no caminho do que uma no rim, beba água.**" "**Você não é um cacto, vai lá beber água!**") }}
{{ $listemote := (cslice "<:wt1:760214268997140510>" "<:wt2:760214268880093204>" "<:wt3:768036987679801384>" "<:wt4:768106746794999828>" "<:wt5:769036898013413376>" "<:wt6:769036936856993793>" "<:wt7:771009060162109440>") }}  

{{ $pollCh := (index (shuffle $listCh) 0) }}
{{ $pollMsg := (index (shuffle $listmsg) 0) }}
{{ $pollEmote := (index (shuffle $listemote) 0) }}

{{ $msg := (joinStr "" ($pollEmote) " | " ($pollMsg)) }}
{{ $msgID:=sendMessageNoEscapeRetID $pollCh $msg }}
{{deleteMessage $pollCh $msgID 30}}