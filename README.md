## OPSWAT Coding Assignment
### Xinchi He (hexinchi@gmail.com)

Thanks for reviewing my code. The code is written in Go. Please follow the instructions to run the code.

The following procedure assumes you are running with the Linux environment, specifically the Debian-based distribution

1. Install Go
```
$ sudo apt install golang
```

2. Clone the repository
```
$ git clone https://github.com/xinchi-he/opswat
```


3. Insert your API KEY by open the lookup.go file in a text editor, replace your key against line 32.

4. Install 3rd-party library for JSON parsing
```
$ cd opswat
$ go get github.com/buger/jsonparser
```

5. Run the code
```
$ go run lookup.go yourfile
```
There are several testing files I have used in the testing folder, which allows you see the report immediately, this is because it takes time for the cloud to run the newly uploaded file against different engines.

Thank you.

Here is a sample output:
```
[~/opswat]$ go run lookup.go testing/MetaDefender-Secure-Networks-Datasheet-RGB.pdf

record found!
filename: MetaDefender-Secure-Networks-Datasheet-RGB.pdf
overall_status: No Threat Detected

engine: Zillya!
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T15:41:00.000Z

engine: Xvirus Personal Guard
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T03:47:00.000Z

engine: Windows Defender
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T15:29:45.000Z

engine: VirusBlokAda
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T14:47:00.000Z

engine: TrendMicro House Call
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-12T20:38:00.000Z

engine: TrendMicro
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-12T20:23:00.000Z

engine: Total Defense
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T00:00:00.000Z

engine: ThreatTrack
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T08:30:00.000Z

engine: TACHYON
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T05:00:00.000Z

engine: Sophos
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T02:54:00.000Z

engine: SUPERAntiSpyware
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-10T20:55:00.000Z

engine: Quick Heal
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T06:21:00.000Z

engine: Preventon
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T15:15:00.000Z

engine: NANOAV
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T17:48:00.000Z

engine: McAfee
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T00:00:00.000Z

engine: Kaspersky
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T13:20:00.000Z

engine: K7
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T16:41:00.000Z

engine: Jiangmin
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T07:02:00.000Z

engine: Ikarus
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T19:22:21.000Z

engine: Huorong
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T08:27:00.000Z

engine: Hauri
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-15T00:00:00.000Z

engine: Fortinet
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T00:00:00.000Z

engine: Filseclab
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T23:17:00.000Z

engine: F-prot
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T23:22:00.000Z

engine: Emsisoft
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T22:37:00.000Z

engine: ESET
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T00:00:00.000Z

engine: Cyren
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-15T01:03:00.000Z

engine: CrowdStrike Falcon ML
threat_found: CLEAN 
scan_result: 23
def_time: 2020-01-15T00:00:00.000Z

engine: Comodo
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-15T00:20:36.000Z

engine: ClamAV
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T13:34:00.000Z

engine: ByteHero
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T00:00:00.000Z

engine: BitDefender
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T22:04:00.000Z

engine: Avira
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T00:00:00.000Z

engine: Antiy
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T15:47:00.000Z

engine: Ahnlab
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-15T00:00:00.000Z

engine: Agnitum
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-12T23:00:00.000Z

engine: AegisLab
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T07:45:00.000Z

engine: Vir.IT eXplorer
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-14T14:46:00.000Z

engine: Vir.IT ML
threat_found: CLEAN 
scan_result: 0
def_time: 2020-01-13T14:46:00.000Z

```
