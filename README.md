# GO MERCADO PAGO

Este modulo es como un mini SDK de mercadopago hecho con go, al cual agregue funcionalidades
para facilitar el uso de la api.

## Prerequisitos
Version de go >= 1.19

## Principales estructuras predefinidas
1. Configuration: estructura que almacena los datos de la configuracion. Ej: accessToken.
    - ubicacion: package config
2. Preference: estructura que almacena los datos de las preferencias a crear.
    - ubicacion: package preference.
3. ConfigurationPreference: estructura que almacena los datos por defecto de la preferencia antes de que esta sea creada.
    - ubicacion: package preference.
4. ResponsePreference: estructura que almacena los datos de la preferencia creada.
    - ubicacion: package preference.
5. Order: estructura que almacena los datos de las ordenes a crear.
    - ubicacion: package order.
6. ResponseOrder: estructura que almacena los datos de la orden creada.
    - ubicacion: package order.

## Funcionalidades

### Package Config
- func Set(c Configuration).
    - Parametros:
        - Configuration: estructura que almacena datos de configuracion. Ej:accessToken.
    - Descripcion: recibe un parametro de tipo Configuration y guarda la configuracion en una variable del mismo tipo.

### Package Preference
- func Create(configurationPreference *ConfigurationPreference) (ResponsePreference, error).
    - Paramentros:
        - *ConfigurationPreference: estructura que almacena los datos por defecto de la preferencia antes de que esta se cree.
    - Descripcion: crea la preferencia y retorna el resultado, tambien retorna un error en caso de haber uno.
- func setConfigurationPreference(p *Preference, configurationPreference *ConfigurationPreference)
    - Parametros:
        - *Preference: datos de la preferencia.
        - *ConfigurationPreference: estructura que almacena los datos por defecto de la preferencia antes de que esta se cree.
    - Descripcion: configura los datos que quiero que esten por defecto antes de crear la preferencia, por ejemplo los "installments"(cantidad de cuotas permitidas), que por ahora es lo unico agregado.

## Package Order
- func Create() (ResponseOrder, error).
    - Parametros: ninguno.
    - Descripcion: Crea la orden y retorna un ResponseOrder y un error, el primero es la orden creada y el segundo contiene el error en caso de que haya uno.