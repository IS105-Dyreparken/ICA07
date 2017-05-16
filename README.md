# ICA07


Kilde for koden: https://www.socketloop.com/tutorials/golang-udp-client-server-read-write-example



UDP-server

For å kjøre filen skriver man i terminal: go run udpserver.go


Serveren er en ”listener” og vil ”høre” etter UDP-clienter på din localhost, dette betyr at serveren prøver å fange opp infomasjon fra client.

# beskrivelse av koden

"Func handleUDPConnection" er en funksjon som søker etter UDP-client. Videre ser vi at en buffer er definert. En buffer er en 


En buffer er en midlertidlig lagringsplass med variabel størrelse i byte, og innholder lese- og skrive-metoder. Nullverdien for en buffer er at den er tom og klar til bruk.

På linje 17 og 18 har vi funksjoner som skriver ut adressen den har kobelt seg til, og en melding eller data som den motar fra cliente i form av type String. Rett under har vi en if-settning som gir tilbakemelidnger om evt. feil.

på linje 25 definerer vi verdien til variablen "message" i dette tilfellet er det en "hemmelig melding". I likhet med if-setningen ovenfor vil denne feilmelidngen også gi tilbakemelding på evt- feil.

Func main
Setter opp UDP-serveren med hostname og portnummer, og variabelen "service" er definert som hostName + portNum. Etterfulgt av dette er det enda en error håndtering.

Videre setter man opp ”listener” hvor den leter etter inkommende UDP-kobling.

Det siste er en for-each løkke, hvor den venter på at UDP-clienten som kobles til.


UDP-client
For å kjøre denne skriver man: go run udpclient.go

Importerer pakkene log, net og fmt.

Func main()

Inneholder spesifikasjonene til clienten, hostnamer og portnummer.
Linje 19, her er hvor den prøver å koble seg til UDP-serveren.


Linje 32, her sender den en melding til UDP-serveren.

42 og videre: her mottar den informasjon og meldinger fra UDP-serveren


