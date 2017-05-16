Ica 07   oppgave 2


1)	Lag en TCP server og klient.
TCP server

Vi har en TCP enkel TCP server som klarer å kommunisere med hver andre ved å skrive en «melding».

I koden begynner vi med å definere porten og om det skal være UDP eller TCP.
Så skal den godta koblingen til porten.
Så kjører vi en for løkke som ved hjelp av bufio for å «lytte» til meldingen som skal skriver ut ved hjelp av en fmt.print ln som skal da printe ut meldingen som er sendt fra klienten.

TCP klient

TCP klient koden er ganske enkel. Vi har forsøker å koble til socket, der ifra har vi en for løkke som skal lese inputen ved hjelp av os.stdin pakken og printer ut den «teksten vi skal sende». Når det er gjort så sender vi teksten til socket og venter på svar. Til slutt har vi en print som vier message fra server.

b) vi fikk dessverre ikke studert dette i wireshark. Vi hadde ingen Linux eller Mac tilgjengelig. Vi forsøkte å få dette til på Windows, men det var for komplisert.
i)  
både TCP og UDP er protokoller som blir brukt for å sende bits av data. Begge er bygget opp på «top of the internet protocol» med andre ord, om man sender packet via TCP eller UDP, vil den packeten bli sendt til en ip adresse. Disse packets er behandlet likt.
Forskjellen mellom TCP og UDP er at TCP garanterer at mottakeren mottar pakkene I rekkefølge ved å nummere dem. Mottakeren sender meldinger tilbake til avsenderen som sier at den mottok meldingene. Hvis avsenderen ikke får riktig svar, vil den sende pakkene igjen for å sikre at mottakeren mottok dem. Pakker kontrolleres også for feil. Pakker som er sendt med TCP blir sporet slik at ingen data går tap teller ødelagt I transit. Dette er grunnen til at nedlastninger av filen ikke blir ødelagt, selv om det er en nettverkshikke.
Når man bruker UDP, sendes pakker bare til mottakeren. Avsenderen venter ikke på at mottakeren mottok pakken. Det vil bare fortsette å sende de neste pakkene.  Hvis du er mottakeren, og du savner noen UDP pakker, så dårlig du kan ikke be disse pakkene igjen. Det er ingen garanti for at du får alle pakkene, og det er ingen måte å spørre om en pakke igjen, hvis du savner det.

ii)  Hvor stor kan en TCP pakke være?
Samme som UDP, er maks størrelsen på en TCP pakke 65535 bytes. Ulikt UDP, så er størrelsen som blir brukt langt mindre. MTU (Maximum Transmission Unit), som blir brukt av TCP for å kontrollere størrelse, forsikrer at pakkestørrelsen ikke er over 1500 bytes.

iii) hva er fragmentering, hvorfor oppstår det og hvordan håndterer man det?
Forkortelser:
MTU = maximum transmission unit
MSS = maximum segment size

Når en ip datagram er for stor for maximum transmission unit (MTU), av den underliggende datalinklagsteknologien som brukes til neste «leg» av reisen, det må bli fragmentert før den kan bli sent tvers av nettverket. Den høyere lagsmeddelelsen som skal sendes, sendes ikke i et enkelt IP datagram, men heller brutt ned i stykker kalt fragmenter som sendes separat. I noen tilfeller må fragmentene selv måtte bli fragmentert ytterligere.

Fragmenterings Problemer og bekymringer

Fragmentering er nødvendig for å implementere et nettverkslagsinternett som er uavhengig av lavere lagdetaljer, men introduserer betydelig kompleksitet til IP. Vi må huske at ip er en upålitelig forbindelsesløs protokoll. Ip datagrammer kan ta noen av flere ruter på vei fra kilden til destinasjonen, og noen kna ikke engang gjøre det til målet I det hele tatt. Når vi fragmenterer en melding, lager vi et enkelt datagram til mange, som kan introdusere flere nye problemer.

Håndtering:

Maksimal segmentstørrelse (MSS) er den største mengden data, spesifisert i byte, som TCP er villig til å motta i et enkelt segment. For best ytelse, bør MSS settes liten nok til å unngå IP-fragmentering, noe som kan føre til tap av pakker og overdrevne retransmisjoner. For å forsøke å oppnå dette, blir MSS vanligvis annonsert ved hver side ved hjelp av MSS-alternativet når TCP-tilkoblingen er etablert, i så fall er den avledet fra den maksimale transmisjonsenhetens (MTU) størrelse på datalinklaget av nettene som Avsenderen og mottakeren er direkte festet. Videre kan TCP-sendere bruke MTU-sti til å utlede minimum MTU langs nettverksbanen mellom avsender og mottaker, og bruk dette for å dynamisk justere MSS for å unngå IP-fragmentering i nettverket.



iv) I hvilke brukerscenarier bruker man UDP og i hvilke TCP?


UDP:
UDP brukes når hastighet er ønskelig og feilkorreksjon er ikke nødvendig. UDP brukes for eksempel ofte til live-streams og online games.
For eksempel, la oss si at du ser på en live stream. Live streams sendes ofte med UDP I stedet for TCP. Serveren sender bare en constant stream av UDP pakker til datamaskiner som ser på. Hvis du mister forbindelsen I noen sekunder, fryser videoen et øyeblikk og hopper deretter til den nærværende delen av kringkastingen, hopper over brikkene du savnet. Hvis du opplever mindre tap, kan videoen eller lyden bli dårligere et øyeblikk ettersom videoen fortsetter å spilles uten de manglende dataene.

TCP:
Som vi allerede vet TCP er annerledes når det gjelder sending av data, den vil gi tilbake melding til mottareken om meldingen er mottatt og alltid sjekker for error. Eksempler hvor TCP kan brukes er nettlesing, epost og filoverføringer er vanlige applikasjoner som bruker TCP.
