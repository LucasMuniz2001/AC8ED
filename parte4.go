package main

import "fmt"

func main() {
    felizNatal := map[string]string{
        "brasil":          "Feliz Natal!",
        "alemanha":        "Frohliche Weihnachten!",
        "austria":         "Frohe Weihnacht!",
        "coreia":          "Chuk Sung Tan!",
        "espanha":         "Feliz Navidad!",
        "grecia":          "Kala Christougena!",
        "estados-unidos":  "Merry Christmas!",
        "inglaterra":      "Merry Christmas!",
        "australia":       "Merry Christmas!",
        "portugal":        "Feliz Natal!",
        "suecia":          "God Jul!",
        "turquia":         "Mutlu Noeller",
        "argentina":       "Feliz Navidad!",
        "chile":           "Feliz Navidad!",
        "mexico":          "Feliz Navidad!",
        "antardida":       "Merry Christmas!",
        "canada":          "Merry Christmas!",
        "irlanda":         "Nollaig Shona Dhuit!",
        "belgica":         "Zalig Kerstfeest!",
        "italia":          "Buon Natale!",
        "libia":           "Buon Natale!",
        "siria":           "Milad Mubarak!",
        "marrocos":        "Milad Mubarak!",
        "japao":           "Merii Kurisumasu!",
    }

    paises := []string{"uri-online-judge", "alemanha", "brasil", "austria"}

    for _, pais := range paises {
        frase, encontrado := felizNatal[pais]
        if encontrado {
            fmt.Println(frase)
        } else {
            fmt.Println("--- NOT FOUND ---")
        }
    }
}
