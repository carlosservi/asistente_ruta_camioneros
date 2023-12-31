# Metodologías y Herramientas para Test

## Qué son los test unitarios y para que los necesitamos
Las pruebas unitarias consisten en aislar una parte del código y comprobar que funciona como esperamos. Son pequeños tests que validan el comportamiento y la lógica de una función o un trozo de código.

Son muy necesarios para poder comprobar que lo que hemos hecho o vamos a hacer es correcto.

### Beneficios
- Demostrar que la lógica del código está correcto y que funcionará en todos los casos.
- Aumentar la legibilidad del código.
- Sirven como documentación del proyecto.
- Suelen ser muy rápidos y puedes realizar muchos en muy poco tiempo con las herramientas adecuadas.

## Metodologías
Existen dos metodologías principales para los test unitarios que son TDD (Test Driven Development) y BDD (Behavior Driven Development).

También es importante diferenciar que se pueden crear los test antes de hacer la funcionalidad o después.

### TDD
Trata de desarrollar test unitarios a los que someter funciones concretas de un software para probar su lógica y funcionamiento. El objetivo es conseguir un código limpio y que funcione correctamente.

### BDD
Es una evolución de TDD pero que mejora la comunicación entre equipos técnicos y no técnicos. De manera que se centran en el usuario y el comportamiento del sistema, a diferencia de TDD que se centra en funcionalidades. Las pruebas escritas en BDD siguen un formato llamado "Gherkin" que usa palabras clave como Given, When y Then.

#### Ventajas
- Se describen detalles de cómo se debe comportar la función concreta a desarrollar y de ésta forma será comprensible por todos.
- Mejora la comunicación entre desarrolladores, testers, usuarios y la dirección.
- El enfoque de definición ayuda a una aceptación común de las funcionalidades previamente al desarrollo.

### Conclusión
Tras conocer las ventajas y los tipos de metodologías y pensando siempre en el cliente, la metodología BDD es mucho mejor ya que ayuda mucho a la comunicación con el cliente y a la comprensión de éste acerca del proyecto y sus funcionalidades.

## Criterios para elegir las Herramientas de Test
- Compatibilidad con el lenguaje de programación.
- Estandarización y Buenas prácticas del lenguaje
- Comunidad y soporte
- Rendimiento
- Flexibilidad y extensibilidad

## Herramientas para Test en Golang

- Aserciones.- Golang no incluye su propia librería de aserciones, el paquete de test estándar "testing", incluye algunas funciones básicas como 'testing.Errorf', pero nada más. 
Se puede añadir alguna libreria externa, tenemos algunas opciones:
    - Testify: es una biblioteca de aserciones que proporciona un conjunto completo de funciones para afirmar y verificar condiciones. Además está diseñado específicamente para trabajar con "testing" de Go, sigue las buenas prácticas de Go, es muy popular y tiene una comunidad muy activa. Respecto al rendimiento es muy liviano y no agrega una carga.
    - GoConvey: GoConvey es una biblioteca de aserciones con una sintaxis más descriptiva. También incluye una interfaz web que muestra el estado de las pruebas en tiempo real. Diseñado para trabajar con el paquete de testing estándar de Go. Respecto al rendimiento agrega una capa adicional debido a la interfaz web, pero es adecuado para proyectos más grandes.
Siguiendo los proverbios de Golang de buenas prácticas y para mi caso concreto, en principio no voy a usar ninguna libreria externa.

- Test runner.- Golang incluye su propio framework de test en Go packages, siendo también el estándar y proporcionando otras muchas funcionalidades.

- Herramienta CLI para ejecución.- Como no puede ser de otra manera Golang también se encarga automáticamente de ésto. A través del comando ```go test``` ejecuta automáticamente todos los ficheros de test que tengan un formato y sean puestos de una manera concreta en el proyecto. A continuación detallo el formato y la estructura para que se ejecuten correctamente con Golang.

### Formato y estructura para los test en Golang
- Estar en la carpeta test del proyecto
- Fichero que termine en _test.go
- Todos los nombres de funciones que comienzan con Test, aunque no es obligatorio la T en mayúscula pero es buena práctica.

