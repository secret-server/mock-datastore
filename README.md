# mock-datastore
The  mock-datastore 

- The mock data store contains interface points to mock data. 
- The project provides mock data and code examples. 
- Mock data is in YAML format
- The user of the mock server project creates their own set of data based on their specific needs.
- The rationale for a separate project is to help isolate the data from the server, allowing independent management and project improvements over time.

## Users
```
Kevin Kelche:
    id: 1
    name: KevinKelche
    display_name: Kevin Kelche
    email: kevin.kelche@example.com
    password: bleach.out.34
    roles:
        - 0
        - 2
Jason Lang:
    id: 2
    name: JasonLang
    display_name: Jason Lang
    email: jason.lang@example.com
    password: windows.frost6
    roles:
        - 1 
```
    

## Roles
```
Admin:
    name: Admin
    id: 0
    created: 2024-06-19T14:15:22Z
    enabled: true
    isSystem: false
Ca Web Admin:
    name: CA Web Admin
    id: 1
    created: 2024-06-19T14:15:22Z
    enabled: true
    isSystem: true
Readonly:
    name: Readonly
    id: 2
    created: 2024-06-19T14:15:22Z
    enabled: true
    isSystem: true
```

## Secrets
```
CaIntermediateKey:
    name: CA Intermediate Passphrase
    id: 90
    fields:
        private-key-passphrase: caInterTest
caDatabase:
    name: CA Database
    id: 50
    fields:
        password: brav0!
        server: 192.168.10.131
        user: monty
caKey:
    name: CA Private Key Passphrase
    id: 10
    fields:
        private-key-passphrase: interbrav0!
caReadonlyDatabase:
    name: CA Readonly Database
    id: 60
    fields:
        password: brav0!
        server: 192.168.10.131
        user: monty
caSmtp:
    name: CA SMPT Service
    id: 70
    fields:
        password: test
        server: 192.168.10.131
        user: admin
```
