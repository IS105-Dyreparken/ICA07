Ica 07   oppgave 2


1)	Lag en tcp server og klient.
TCP server

Vi har en TCP enkel tcp server  som klarer å kommunisere med hver andre ved å skrive en «melding».

I koden begynner vi med å definere porten og om det skal være udp eller tcp.
Så skal den godta koblingen til porten.
Der etter kjører vi en for looker som ved hjelp av bufio for å «lytte» til meldingen som skal skriver ut ved hjelp av en fmt.print ln som skal da printe ut meldingen som er sendt fra klienten.



TCP klient

TCP klient koden er ganske enkel. Vi har forsøker å koble til socket, der i fra har vi en for looker som skal lese inputen ved hjelp av os.stdin pakken og printer ut den «teksten vi skal sende». Når det er gjort så sender vi teksten til socket og venter på svar. Til slutt har vi en print som vier message fra server.
