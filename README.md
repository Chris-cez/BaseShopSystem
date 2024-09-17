# Base Shop System
Versão básica e open-source do código usado na criação de um sistema de uma loja real. A versão completa com os dados especificos do projeto será criada e vendida de forma privada. Se tiver interesse na versão paga, tratar com o criador do repositorio original.

## Banco de Dados
Versão pensada para o ==GORM==.

**Product**
| Name | Type |
| :---: | :---: |
| Id | string | ==Primary key==
| Name | string |
| Price | float32 |
| NCM | string |
| UM | string |
| Class | uint8 | ==Foreign key==
| Description | string |

**Client**
| Name | Type |
| :---: | :---: |
| Id | string | ==Primary key==
| Name | string |
| CPF | int32 | 
| Address | uint16 | ==Foreign key==
