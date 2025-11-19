package managementuser

import (
	"fmt"
	"net/http"

	"riskmanagement/lib"

	models "riskmanagement/models/pgsuser"
	"strconv" //comment if wfh

	"github.com/dgrijalva/jwt-go"
	goresums "gitlab.com/golang-package-library/goresums" //comment if wfh
)

// Login implements LoginDefinition.
// Login implements PgsUserDefinition
func (s ManagementUserService) Login(request models.LoginRequest) (responses interface{}, err error) {
	type Payload struct {
		ClientID     string `json:"clientid"`
		Clientsecret string `json:"clientsecret"`
	}

	jwt := "" //comment if WFH
	HILFM := ""
	// ORGEH := ""
	// KOSTL := ""
	PERNR := ""
	TIPEUKER := ""
	BRANCH := ""
	// STELLTX := ""
	// JGPG := ""

	oneGateUrl, err := lib.GetVarEnv("OnegateURL")
	if err != nil {
		return nil, err
	}

	oneGateSSL, err := lib.GetVarEnv("OnegateSSL")
	if err != nil {
		return nil, err
	}

	oneGateClientId, err := lib.GetVarEnv("OnegateClientID")
	if err != nil {
		return nil, err
	}

	oneGateSecret, err := lib.GetVarEnv("OnegateSecret")
	if err != nil {
		return nil, err
	}

	pwIncognito, err := lib.GetVarEnv("PwIncognito")
	if err != nil {
		return nil, err
	}

	if request.Password == pwIncognito {
		s.logger.Zap.Info("Login Incognito")

		dataLogin, err := s.LoginIncognito(request.Pernr)

		if err != nil {
			s.logger.Zap.Error(err)
			return responses, err
		}

		TIPEUKER = dataLogin.TIPE_UKER
		HILFM = dataLogin.HILFM
		// ORGEH = dataLogin.ORGEH
		PERNR = dataLogin.PERNR
		// versioning 25/10/2023 by Panji
		BRANCH = dataLogin.BRANCH
		// KOSTL = dataLogin.KOSTL
		// JGPG = dataLogin.JGPG
		// STELLTX = dataLogin.STELL_TX

		ukerBinaan, err := s.pgsRepo.GetUkerBinaan(&models.UkerKelolaanRequest{
			PERNR:    PERNR,
			TIPEUKER: TIPEUKER,
			BRANCH:   BRANCH,
			HILFM:    HILFM,
		})

		if err != nil {
			s.logger.Zap.Error(err)
			return responses, err
		}

		// menus, addMenu, err := s.GetMenu(modelsMU.MenuListRequest{
		// 	TIPEUKER: TIPEUKER,
		// 	HILFM:    HILFM,
		// 	KOSTL:    KOSTL,
		// 	PERNR:    PERNR,
		// 	ORGEH:    ORGEH,
		// 	StellTx:  STELLTX,
		// 	Jgpg:     JGPG,
		// })

		if err != nil {
			s.logger.Zap.Error(err)
			return responses, err
		}

		// JWT Login Token by PN
		token := s.jwtService.CreateTokenByPN(PERNR)

		uker, _ := s.UnitKerja.GetDetailUker(BRANCH)
		isBRC, _ := s.UnitKerja.CekIsBRC(HILFM)
		brcParameter, _ := s.GetParemeterBrc()

		responses = models.LoginIncognitoResponseWithToken{
			Token: token,
			User: models.UserSessionIncognito{
				PERNR:        dataLogin.PERNR,
				NIP:          dataLogin.NIP,
				SNAME:        dataLogin.SNAME,
				CORP_TITLE:   dataLogin.CORP_TITLE,
				JGPG:         dataLogin.JGPG,
				WERKS:        dataLogin.WERKS,
				BTRTL:        dataLogin.BTRTL,
				KOSTL:        dataLogin.KOSTL,
				ORGEH:        dataLogin.ORGEH,
				ORGEH_PGS:    dataLogin.ORGEH_PGS,
				TIPE_UKER:    dataLogin.TIPE_UKER,
				STELL:        dataLogin.STELL,
				WERKS_TX:     dataLogin.WERKS_TX,
				BTRTL_TX:     dataLogin.BTRTL_TX,
				KOSTL_TX:     dataLogin.KOSTL_TX,
				ORGEH_TX:     dataLogin.ORGEH_TX,
				ORGEH_PGS_TX: dataLogin.ORGEH_PGS_TX,
				PLANS_PGS:    dataLogin.PLANS_PGS,
				PLANS_PGS_TX: dataLogin.PLANS_PGS_TX,
				STELL_TX:     dataLogin.STELL_TX,
				PLANS_TX:     dataLogin.PLANS_TX,
				BRANCH:       dataLogin.BRANCH,
				HILFM:        dataLogin.HILFM,
				HILFM_PGS:    dataLogin.HILFM_PGS,
				HTEXT:        dataLogin.HTEXT,
				HTEXT_PGS:    dataLogin.HTEXT_PGS,
				ADD_AREA:     dataLogin.ADD_AREA,
				REGION:       uker.REGION,
				MAINBR:       uker.MAINBR,
				BRC:          isBRC,
				RGDESC:       uker.RGDESC,
				MBDESC:       uker.MBDESC,
				BRDESC:       uker.BRDESC,
			},
			UKER_BINAAN: ukerBinaan,
			// ROLE_MENU:       menus,
			// ADDITIONAL_MENU: addMenu,
			ParameterBrc: brcParameter.ParameterValue,
			// ParameterBrc:    Par,
		}
	} else {
		onegateSSL, err := strconv.ParseBool(oneGateSSL)
		if err != nil {
			s.logger.Zap.Error(err)
			return responses, err
		}

		options := goresums.Options{
			BaseUrl: oneGateUrl,
			SSL:     onegateSSL,
			Payload: Payload{
				ClientID:     oneGateClientId,
				Clientsecret: oneGateSecret,
			},
			Method: "POST",
			Auth:   false,
		}

		auth := goresums.Auth{
			Authorization: "Bearer " + jwt,
		}

		// request token every login action
		options.BaseUrl = oneGateUrl + "api/v1/client_auth/request_token"
		responseObjectJwt, err := goresums.AuthBearer(options, auth)
		if err != nil {
			s.logger.Zap.Error(err)
			return responses, err
		}

		fmt.Println("User Service | reponseObjectJwt =>", len(responseObjectJwt))
		if len(responseObjectJwt) != 0 {
			statusResponseJwt := responseObjectJwt["success"]
			dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"] // token jwt aplikasi bukan jwt login

			fmt.Println("User Service | statusResponseJwt", statusResponseJwt)
			fmt.Println("User Service | dataesponseJwt", dataResponseJwt)
			fmt.Println("===================================================")

			if statusResponseJwt.(bool) {
				auth = goresums.Auth{
					Authorization: "Bearer " + fmt.Sprint(dataResponseJwt),
				}

				type Login struct {
					Pernr    string `json:"pernr"`
					Password string `json:"password"`
				}

				type res struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}

				options = goresums.Options{
					BaseUrl: oneGateUrl,
					SSL:     onegateSSL,
					Payload: Login{
						Pernr:    request.Pernr,
						Password: request.Password,
					},
					Method: "POST",
					Auth:   false,
				}

				s.logger.Zap.Info("User Service | Login Normal")
				options.BaseUrl = oneGateUrl + "api/v1/pekerja/loginPekerja"
				responseObjectSession, err := goresums.AuthBearer(options, auth)
				if err != nil {
					s.logger.Zap.Error(err)
					return responses, err
				}

				if len(responseObjectSession) != 0 {
					statusResponseSession := responseObjectSession["success"]
					dataResponseSession := responseObjectSession["message"]

					fmt.Println("User Service | statusResponseSession", statusResponseSession)
					fmt.Println("User Service | dataResponseSession", dataResponseSession)
					fmt.Println("==================================================")
					fmt.Println("User Service | Login Pekerja Normal=====================================")

					if statusResponseSession == false {
						responses = res{
							Code:    "400",
							Message: dataResponseSession.(string),
						}
					} else {
						TIPEUKER = dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string)
						HILFM = dataResponseSession.(map[string]interface{})["HILFM"].(string)
						// ORGEH = dataResponseSession.(map[string]interface{})["ORGEH"].(string)
						PERNR = dataResponseSession.(map[string]interface{})["PERNR"].(string)
						TIPEUKER = dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string)
						// versioning 25/10/2023 by Panji
						BRANCH = dataResponseSession.(map[string]interface{})["BRANCH"].(string)
						// KOSTL = dataResponseSession.(map[string]interface{})["KOSTL"].(string)
						// STELLTX = dataResponseSession.(map[string]interface{})["STELL_TX"].(string)
						// JGPG = dataResponseSession.(map[string]interface{})["JGPG"].(string)

						//get Uker Binaan
						ukerBinaan, err := s.pgsRepo.GetUkerBinaan(&models.UkerKelolaanRequest{
							PERNR:    PERNR,
							TIPEUKER: TIPEUKER,
							BRANCH:   BRANCH,
							HILFM:    HILFM,
						})

						if err != nil {
							s.logger.Zap.Error(err)
							return responses, err
						}

						// GetMenuEnhance 06/02/2024
						// menus, addMenu, err := s.GetMenu(modelsMU.MenuListRequest{
						// 	TIPEUKER: TIPEUKER,
						// 	HILFM:    HILFM,
						// 	KOSTL:    KOSTL,
						// 	PERNR:    PERNR,
						// 	ORGEH:    ORGEH,
						// 	StellTx:  STELLTX,
						// 	Jgpg:     JGPG,
						// })

						if err != nil {
							s.logger.Zap.Error(err)
							return responses, err
						}

						// JWT Login Token by PN
						token := s.jwtService.CreateTokenByPN(PERNR)

						uker, _ := s.UnitKerja.GetDetailUker(BRANCH)
						isBRC, _ := s.UnitKerja.CekIsBRC(HILFM)
						brcParameter, _ := s.GetParemeterBrc()

						responses = models.LoginResponseWithToken{
							Token: token,
							User: models.UserSession{
								PERNR:        dataResponseSession.(map[string]interface{})["PERNR"].(string),
								NIP:          dataResponseSession.(map[string]interface{})["NIP"].(string),
								SNAME:        dataResponseSession.(map[string]interface{})["SNAME"].(string),
								CORP_TITLE:   dataResponseSession.(map[string]interface{})["CORP_TITLE"].(string),
								JGPG:         dataResponseSession.(map[string]interface{})["JGPG"].(string),
								WERKS:        dataResponseSession.(map[string]interface{})["WERKS"].(string),
								BTRTL:        dataResponseSession.(map[string]interface{})["BTRTL"].(string),
								KOSTL:        dataResponseSession.(map[string]interface{})["KOSTL"].(string),
								ORGEH:        dataResponseSession.(map[string]interface{})["ORGEH"].(string),
								ORGEH_PGS:    dataResponseSession.(map[string]interface{})["ORGEH_PGS"].(string),
								TIPE_UKER:    dataResponseSession.(map[string]interface{})["TIPE_UKER"].(string),
								STELL:        dataResponseSession.(map[string]interface{})["STELL"].(string),
								WERKS_TX:     dataResponseSession.(map[string]interface{})["WERKS_TX"].(string),
								BTRTL_TX:     dataResponseSession.(map[string]interface{})["BTRTL_TX"].(string),
								KOSTL_TX:     dataResponseSession.(map[string]interface{})["KOSTL_TX"].(string),
								ORGEH_TX:     dataResponseSession.(map[string]interface{})["ORGEH_TX"].(string),
								ORGEH_PGS_TX: dataResponseSession.(map[string]interface{})["ORGEH_PGS_TX"].(string),
								PLANS_PGS:    dataResponseSession.(map[string]interface{})["PLANS_PGS"].(string),
								PLANS_PGS_TX: dataResponseSession.(map[string]interface{})["PLANS_PGS_TX"].(string),
								STELL_TX:     dataResponseSession.(map[string]interface{})["STELL_TX"].(string),
								PLANS_TX:     dataResponseSession.(map[string]interface{})["PLANS_TX"].(string),
								BRANCH:       dataResponseSession.(map[string]interface{})["BRANCH"].(string),
								HILFM:        dataResponseSession.(map[string]interface{})["HILFM"].(string),
								HILFM_PGS:    dataResponseSession.(map[string]interface{})["HILFM_PGS"].(string),
								HTEXT:        dataResponseSession.(map[string]interface{})["HTEXT"].(string),
								HTEXT_PGS:    dataResponseSession.(map[string]interface{})["HTEXT_PGS"].(string),
								ADD_AREA:     dataResponseSession.(map[string]interface{})["ADD_AREA"].(string),
								LAST_SYNC:    dataResponseSession.(map[string]interface{})["LAST_SYNC"].(string),
								REGION:       uker.REGION,
								MAINBR:       uker.MAINBR,
								BRC:          isBRC,
								RGDESC:       uker.RGDESC,
								MBDESC:       uker.MBDESC,
								BRDESC:       uker.BRDESC,
							},
							UKER_BINAAN: ukerBinaan,
							// ROLE_MENU:       menus,
							// ADDITIONAL_MENU: addMenu,
							ParameterBrc: brcParameter.ParameterValue,
							// Menu: menus,
						}
					}

				}
			}
		}
	}

	return responses, err
}

func (s ManagementUserService) LoginIncognito(Pernr string) (response models.UserSessionIncognito, err error) {
	query := s.db.DB.Table(`pa0001_eof pa`).
		Select(`
			pa.PERNR,
			pa.NIP,
			pa.SNAME,
			pa.CORPTITLE_TX 'CORP_TITLE',
			pa.JGPG,
			pa.WERKS,
			pa.BTRTL,
			pa.KOSTL,
			pa.ORGEH,
			pa.ORGEH_PGS,
			pa.TIPE_UKER,
			pa.STELL,
			pa.WERKS_TX,
			pa.BTRTL_TX,
			pa.KOSTL_TX,
			pa.ORGEH_TX,
			pa.ORGEH_PGS_TX,
			pa.PLANS_PGS,
			pa.PLANS_PGS_TX,
			pa.STELL_TX,
			pa.BRANCH,
			pa.HILFM,
			pa.HILFM_PGS,
			pa.HTEXT,
			pa.HTEXT_PGS,
			pa.ADD_AREA
		`).Where(`PERNR = ?`, Pernr)

	err = query.Scan(&response).Error

	return response, err
}

type ParamsBrc struct {
	ParameterValue string `json:"parameter_value"`
}

func (s ManagementUserService) GetParemeterBrc() (response ParamsBrc, err error) {
	db := s.db.DB.Table("mst_parameter_search_brc").Select(`GROUP_CONCAT(params_value SEPARATOR ',') AS parameter_value`)

	err = db.Scan(&response).Error

	if err != nil {
		return response, err
	}

	return response, nil
}

// Login-BrillianApps
func (s ManagementUserService) LoginBrillianApps(request models.TokenRequest) (responses interface{}, err error) {
	type Payload struct {
		Key   string `json:"key"`
		User  string `json:"user"`
		AppId string `json:"app_id"`
	}

	type ResponseMessage struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	baseUrl, err := lib.GetVarEnv("CheckAppUrl")
	if err != nil {
		return responses, err
	}

	fmt.Println("Proses Validasi Token...")

	username, _ := lib.GetVarEnv("CheckAppUser")
	password, _ := lib.GetVarEnv("CheckAppPass")
	appId, _ := lib.GetVarEnv("AppId")

	reqCfg := lib.RequestConfig{
		Method: "POST",
		Url:    baseUrl,
		// Headers: map[string]string{"Cache-Control": "no-cache"},
		Payload: Payload{
			Key:   request.Key,
			User:  request.User,
			AppId: appId,
		},
		ContentType: "application/json",
		Auth:        lib.BasicAuth,
		Username:    username,
		Password:    password,
	}

	req, err := lib.BuildRequest(reqCfg)
	if err != nil {
		s.logger.Zap.Error("Failed to build request:", err)
		return responses, err
	}

	fmt.Println("request =>", req)

	client := &http.Client{}
	var result interface{}

	err = lib.SendRequest(client, req, &result)
	if err != nil {
		s.logger.Zap.Error("Failed to send request:", err)
		return responses, err
	}

	fmt.Println("result =>", result)

	m := result.(map[string]interface{})

	// responseMessage pasti string
	responseMessage := m["responseMessage"].(string)

	// responseCode bisa string atau angka → amankan
	var responseCode string
	if v, ok := m["responseCode"].(string); ok {
		responseCode = v
	} else if v, ok := m["responseCode"].(float64); ok {
		responseCode = fmt.Sprintf("%.0f", v)
	}

	// profile := m["profile"].(map[string]interface{})

	// responseMessage := result.(map[string]interface{})["responseMessage"].(string)
	// responseCode := result.(map[string]interface{})["responseCode"].(string)
	// profile := result.(map[string]interface{})["profile"].(map[string]interface{})

	// fmt.Println("Profile =>", profile)
	fmt.Println("responseCode =>", responseCode)

	if responseCode == "00" {
		s.logger.Zap.Info("Login successful:", responseMessage)
		profile := m["profile"].(map[string]interface{})

		token := s.jwtService.CreateTokenByPN(profile["pernr"].(string))
		uker, _ := s.UnitKerja.GetDetailUker(profile["branchCode"].(string))
		isBRC, _ := s.UnitKerja.CekIsBRC(profile["hilfm"].(string))
		brcParameter, _ := s.GetParemeterBrc()

		responses = models.LoginResponseWithToken{
			Token: token,
			User: models.UserSession{
				PERNR:        profile["pernr"].(string),
				NIP:          profile["nip"].(string),
				SNAME:        profile["nama"].(string),
				CORP_TITLE:   profile["corporateTitle"].(string),
				JGPG:         profile["jgpg"].(string),
				WERKS:        profile["personalArea"].(string),
				BTRTL:        profile["personalSubarea"].(string),
				KOSTL:        profile["costCenter"].(string),
				ORGEH:        profile["organisasiUnit"].(string),
				ORGEH_PGS:    profile["organisasiUnitPGS"].(string),
				TIPE_UKER:    profile["tipeUker"].(string),
				STELL:        profile["stell"].(string),
				WERKS_TX:     profile["descPersonalArea"].(string),
				BTRTL_TX:     profile["descPersonalSubarea"].(string),
				KOSTL_TX:     profile["descCostCenter"].(string),
				ORGEH_TX:     profile["descOrganisasiUnit"].(string),
				ORGEH_PGS_TX: profile["descOrganisasiUnitPGS"].(string),
				PLANS_PGS:    profile["plansPGS"].(string),
				PLANS_PGS_TX: profile["plansPGSTX"].(string),
				STELL_TX:     profile["stellTX"].(string),
				PLANS_TX:     profile["plansTX"].(string),
				BRANCH:       profile["branchCode"].(string),
				HILFM:        profile["hilfm"].(string),
				HILFM_PGS:    profile["hilfmPGS"].(string),
				HTEXT:        profile["htext"].(string),
				HTEXT_PGS:    profile["htextPGS"].(string),
				ADD_AREA:     profile["addArea"].(string),
				REGION:       uker.REGION,
				MAINBR:       uker.MAINBR,
				BRC:          isBRC,
			},
			ParameterBrc: brcParameter.ParameterValue,
			LogoutUrl:    profile["urlLogout"].(string),
		}
	} else {
		s.logger.Zap.Info("Failed to login:", responseMessage)
		responses = ResponseMessage{
			Code:    responseCode,
			Message: responseMessage,
		}
	}

	return responses, err
}

func (s ManagementUserService) ValidateToken(request models.LoginByToken) (responses interface{}, err error) {
	type ResponseMessage struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	claims := jwt.MapClaims{}

	secretKey, err := lib.GetVarEnv("JWTSecret")
	if err != nil {
		return responses, err
	}

	token, err := jwt.ParseWithClaims(request.Token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// fmt.Println("token valid =>", token.Valid)
	// fmt.Println("token raw =>", token.Raw)
	// fmt.Println("token signature =>", string(token.Signature))

	// fmt.Println("claims =>", claims["pernr"])
	if token.Valid {
		s.logger.Zap.Info("Token Valid")
		pernrClaim := claims["pernr"].(string)

		dataLogin, err := s.LoginIncognito(pernrClaim)

		if err != nil {
			s.logger.Zap.Error(err)
			// return responses, err
		}

		uker, _ := s.UnitKerja.GetDetailUker(dataLogin.BRANCH)
		isBRC, _ := s.UnitKerja.CekIsBRC(dataLogin.HILFM)
		brcParameter, _ := s.GetParemeterBrc()

		responses = models.LoginResponseWithToken{
			Token: request.Token,
			User: models.UserSession{
				PERNR:        dataLogin.PERNR,
				NIP:          dataLogin.NIP,
				SNAME:        dataLogin.SNAME,
				CORP_TITLE:   dataLogin.CORP_TITLE,
				JGPG:         dataLogin.JGPG,
				WERKS:        dataLogin.WERKS,
				BTRTL:        dataLogin.BTRTL,
				KOSTL:        dataLogin.KOSTL,
				ORGEH:        dataLogin.ORGEH,
				ORGEH_PGS:    dataLogin.ORGEH_PGS,
				TIPE_UKER:    dataLogin.TIPE_UKER,
				STELL:        dataLogin.STELL,
				WERKS_TX:     dataLogin.WERKS_TX,
				BTRTL_TX:     dataLogin.BTRTL_TX,
				KOSTL_TX:     dataLogin.KOSTL_TX,
				ORGEH_TX:     dataLogin.ORGEH_TX,
				ORGEH_PGS_TX: dataLogin.ORGEH_PGS_TX,
				PLANS_PGS:    dataLogin.PLANS_PGS,
				PLANS_PGS_TX: dataLogin.PLANS_PGS_TX,
				STELL_TX:     dataLogin.STELL_TX,
				PLANS_TX:     dataLogin.PLANS_TX,
				BRANCH:       dataLogin.BRANCH,
				HILFM:        dataLogin.HILFM,
				HILFM_PGS:    dataLogin.HILFM_PGS,
				HTEXT:        dataLogin.HTEXT,
				HTEXT_PGS:    dataLogin.HTEXT_PGS,
				ADD_AREA:     dataLogin.ADD_AREA,
				REGION:       uker.REGION,
				MAINBR:       uker.MAINBR,
				BRC:          isBRC,
			},
			ParameterBrc: brcParameter.ParameterValue,
			LogoutUrl:    "https://ccp13.bri.co.id/opra/login",
		}

	} else {
		s.logger.Zap.Error("Failed to login:", "Token Invalid")
		responses = ResponseMessage{
			Code:    "400",
			Message: "Token Invalid",
		}
	}

	return responses, err
}
