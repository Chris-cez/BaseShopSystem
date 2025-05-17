#!/bin/sh

http POST 'localhost:8080/api/company' \
    name='Empresa' \
    cnpj='00.000.000/0000-00' \
    inscricao_estadual='000000' \
    password='Senha' \
    address_id='';