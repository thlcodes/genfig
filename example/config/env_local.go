// Code generated by genfig (config built by merging 'default.yml' and '.env.local') on 2019-07-03T22:39:42+02:00; DO NOT EDIT.

package config

func init() {
  Envs.Local = Config{
    Secrets: []string{"ChuckNorriscanwinagameofConnectFourinonlythreemoves"},
    Randomizer: ConfigRandomizer{
      Threshold: 0.12345,
    },
    LongDesc: ConfigLongDesc{
      En: "Long description",
      De: "Lange Beschreibung",
    },
    Version: "0.1.0",
    Db: ConfigDb{
      Pass: "norris",
      Uri: "mongdb://localhos:27017/db",
      User: "chuck",
    },
    Apis: ConfigApis{
      Google: ConfigApisGoogle{
        Uri: "google.com",
      },
    },
    Wip: true,
    Project: "genfig",
    Server: ConfigServer{
      Host: "localhost",
      Port: 1212,
    },
  }
}