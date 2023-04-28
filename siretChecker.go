package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//coucou
	OBSMap := make(map[string]bool)
	OBSFileName := os.Args[2]
	OBSFile, err := os.Open(OBSFileName)
	if err != nil {
		log.Fatalf("Couldn't open %v\n", err)
	}
	OBSReader := csv.NewReader(bufio.NewReader(OBSFile))
	OBSReader.Comma = ';'

	for i := 0; ; i++ {
		record, err := OBSReader.Read()
		if err == io.EOF || len(record[0]) == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("-> %v\n", record[0])
		OBSMap[record[0]] = true
	}
	//fmt.Printf("%v\n", OBSMap)

	inseeFileName := os.Args[1]
	inseeFile, err := os.Open(inseeFileName)
	if err != nil {
		log.Fatalf("Couldn't open %v\n", err)
	}
	inseeReader := csv.NewReader(bufio.NewReader(inseeFile))
	inseeReader.Comma = ','
	fmt.Println("siren,nic,siret,statutDiffusionEtablissement,dateCreationEtablissement,trancheEffectifsEtablissement,anneeEffectifsEtablissement,activitePrincipaleRegistreMetiersEtablissement,dateDernierTraitementEtablissement,etablissementSiege,nombrePeriodesEtablissement,complementAdresseEtablissement,numeroVoieEtablissement,indiceRepetitionEtablissement,typeVoieEtablissement,libelleVoieEtablissement,codePostalEtablissement,libelleCommuneEtablissement,libelleCommuneEtrangerEtablissement,distributionSpecialeEtablissement,codeCommuneEtablissement,codeCedexEtablissement,libelleCedexEtablissement,codePaysEtrangerEtablissement,libellePaysEtrangerEtablissement,complementAdresse2Etablissement,numeroVoie2Etablissement,indiceRepetition2Etablissement,typeVoie2Etablissement,libelleVoie2Etablissement,codePostal2Etablissement,libelleCommune2Etablissement,libelleCommuneEtranger2Etablissement,distributionSpeciale2Etablissement,codeCommune2Etablissement,codeCedex2Etablissement,libelleCedex2Etablissement,codePaysEtranger2Etablissement,libellePaysEtranger2Etablissement,dateDebut,etatAdministratifEtablissement,enseigne1Etablissement,enseigne2Etablissement,enseigne3Etablissement,denominationUsuelleEtablissement,activitePrincipaleEtablissement,nomenclatureActivitePrincipaleEtablissement,caractereEmployeurEtablissement")
	for i := 0; ; i++ {
		record, err := inseeReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		_, exist := OBSMap[record[2]]
		if exist {
			for j := 0; j < 48; j++ {
				fmt.Printf("%v,", record[j])
			}
			fmt.Println("")
		}
	}
}
