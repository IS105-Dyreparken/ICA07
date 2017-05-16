<<<<<<< HEAD
# ICA07 for gruppe 2, Dyreparken

Består av:  

Dyb, Nikolai Holmen

Hassan, Shiwan

Krossen, Ella

Sandøy, Benjamin AG

Thompson, Sindre  

Yang, Erik
=======
Ica 7 oppgave 1

Analyse av UDP
I vårt eksempel kan vi se at headeren har en lengde på 36 bytes, vi teller generelt bytes fra nettverkslaget og oppover. Hele pakken vår er på 60 bytes. Det vil utgjør 60% av pakken som er den nødvendige prosenten for å transportere meldingen.

I forhold til wikipedia så har vi funnet ut at en UDP pakke har en teoretisk størrelse på 65535 bytes, minus størrelsen på header data. Vi har da 8 bytes for UDP header og 20 bytes for ip header. Da får vi et resultat på 65507 bytes.
Størrelsen på en UDP pakke er 65535 bytes, men begrensninger (-8byte UDP header, - 20 byte IP header) pålagt av IPV4 protokollen gjør den egentlige størrelsen om til 65507.

Vi fikk dessverre ikke tid til å analysere resten på NIC eller TCP fordi vi hadde ikke verken Linux eller MAC tilgjengelig. Vi forsøkte å få det til på Windows, men det var for komplisert for oss.
>>>>>>> Sindre
