# Base Shop System
Versão básica e open-source do código usado na criação de um sistema de uma loja real. A versão completa com os dados especificos do projeto será criada e vendida de forma privada. Se tiver interesse na versão paga, tratar com o criador do repositorio original.

## Banco de Dados
Versão pensada para o **GORM**.

**Product**
| Name | Type |
| :---: | :---: |
| Id | Gorm.model | Primary key
| Name | string |
| Price | float32 |
| NCM | string |
| UM | string |
| Class | uint8 | Foreign key
| Description | string |

**Client**
| Name | Type |
| :---: | :---: |
| Id | Gorm.model | Primary key
| Name | string |
| CPF | int32 | 
| Address | int | Foreign key

**Address**
| Name | Type |
| :---: | :---: |
| Id | Gorm.model | Primary key
| Street | string |
| City | string |
| State | string |

**User**
| Name | Type |
| :---: | :---: |
| Id | Gorm.model | Primary Key
| Name | string |
| CPF | uint32 |
| Id_address | int | Foreign Key

**Class**
| Name | Type |
| :---: | :---: |
| Id | Gorm.model | Primary Key
| Name | string |
| Description | string |

```
type User struct {
	gorm.Model
	Name string
	CPF  uint32
	Id_address uint16
}

type Address struct {
	gorm.Model
	Street string
	City   string
	State  string
}

type Product struct {
	gorm.Model
	Code  string
	Price float32
	Name  string
	NCM   string
	UM    string
	Description string
	Id_Class int
}

type Class struct {
	gorm.Model
	Name string
	Description string
}
```

