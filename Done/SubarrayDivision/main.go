package main


func Birthdays(s []int32, d int32, m int32) int32 {
    // Asegurarse de que m sea positivo y no mayor que la longitud del slice.
    if m <= 0 || int(m) > len(s) {
        return 0
    }
    
    count := int32(0)
    var currentSum int32 = 0

    // Calcular la suma del primer segmento de longitud m.
    for i := 0; i < int(m); i++ {
        currentSum += s[i]
    }
    
    // Si la suma inicial es igual a d, incrementamos el contador.
    if currentSum == d {
        count++
    }
    
    // Deslizar la ventana a través del slice.
    for i := int(m); i < len(s); i++ {
        // Restar el elemento que sale (en la posición i-m)
        currentSum -= s[i-int(m)]
        // Agregar el elemento nuevo que entra en la ventana.
        currentSum += s[i]
        
        if currentSum == d {
            count++
        }
    }
    
    return count
}