{{$listCh:= cslice 456915295940902953 788996947423920128 788997005523288135 788997044650901518 768988994586279936}}
{{$pollCh:=(index (shuffle $listCh) 0)}}

{{$listMsg:= cslice "<:Aviso:789940305171709983> | Personalize seu perfil através dos leilões que rolam no servidor, utilize o comando **-leilão** e fique por dentro de todos os leilões!" "<:Aviso:789940305171709983> | Fique por dentro de tudo que acontece no servidor, adquira o cargo eventos, utilizando o comando **gb.eventos**." ""}}
{{$pollMsg:=(index (shuffle $listMsg) 0)}}

{{$msgID:=sendMessageNoEscapeRetID $pollCh (print $pollMsg)}}
{{deleteMessage $pollCh $msgID 60}}