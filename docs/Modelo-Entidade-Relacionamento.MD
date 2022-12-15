```mermaid
erDiagram
    TECHNOLOGY ||--|{ VERSION : contains
    TECHNOLOGY {
        string name PK
    }
    VERSION {
        string number PK
        string technology_name FK
    }
```