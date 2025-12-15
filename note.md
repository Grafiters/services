UPDATE PER - 25 November 2025
    (use "git add <file>..." to update what will be committed)
    (use "git restore <file>..." to discard changes in working directory)
            modified:   controllers/riskcontrol/riskcontrol_controller.go
            modified:   go.mod
            modified:   go.sum
            modified:   lib/encrypt.go
            modified:   lib/time.go
            modified:   models/riskcontrol/riskcontrol.dto.go
            modified:   models/riskcontrol/riskcontrol.go
            modified:   repository/riskcontrol/riskcontrol_repository.go
            modified:   routes/riskcontrol.go
            modified:   services/auth/jwt_auth_service.go
            modified:   services/riskcontrol/riskcontrol_service.go

    Untracked files:
    (use "git add <file>..." to include in what will be committed)
            dto/
            lib/file.go
            lib/messages.go
            lib/textFormating.go
            models/riskcontrolattribute/
            note.md
    
    penambahan library for export to pdf
        source -> github.com/jung-kurt/gofpdf
        lisencse -> https://github.com/jung-kurt/gofpdf?tab=MIT-1-ov-file

changes code
 controllers/riskcontrol/riskcontrol_controller.go |   7 ++++++-
 dto/arlods.go                                     |  37 +++++++++++++++++++++++++++++++++++++
 lib/httpReq.go                                    |   2 +-
 models/riskcontrol/riskcontrol.dto.go             |   2 +-
 services/arlords/arlords_service.go               |  82 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 services/riskcontrol/riskcontrol_service.go       | 150 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++---------------------------
 services/services.go                              |   2 ++

UPDATE PER - 26 November 2025
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   controllers/riskindicator/riskindicator_controller.go
        modified:   dto/risk_control.go
        modified:   lib/number.go
        modified:   lib/time.go
        modified:   models/riskindicator/riskindicator.dto.go
        modified:   models/riskindicator/riskindicator.go
        modified:   repository/riskindicator/lampiran_indicator_repository.go
        modified:   repository/riskindicator/risk_indicator_repository.go
        modified:   repository/riskissue/risk_issue_repository.go
        modified:   routes/riskindicator.go
        modified:   services/riskcontrol/riskcontrol_service.go
        modified:   services/riskindicator/risk_indicator_service.go
        modified:   services/riskissue/risk_issue_service.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        note.md

UPDATE PER 28 - November 2025
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   controllers/riskindicator/riskindicator_controller.go
        modified:   controllers/riskissue/risk_issue_controller.go
        modified:   dto/arlods.go
        modified:   dto/risk_control.go
        modified:   lib/number.go
        modified:   lib/textFormating.go
        modified:   lib/time.go
        modified:   models/riskindicator/riskindicator.dto.go
        modified:   models/riskindicator/riskindicator.go
        modified:   models/riskissue/risk_issue.dto.go
        modified:   repository/riskindicator/lampiran_indicator_repository.go
        modified:   repository/riskindicator/risk_indicator_repository.go
        modified:   repository/riskissue/risk_issue_repository.go
        modified:   routes/riskindicator.go
        modified:   routes/riskissue.go
        modified:   services/arlords/arlords_service.go
        modified:   services/riskcontrol/riskcontrol_service.go
        modified:   services/riskindicator/risk_indicator_service.go
        modified:   services/riskissue/risk_issue_service.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        dto/arlordbusinessprocess.dto.go
        note.md

Update PER 01 - 12 - 2025
user@RYU MINGW64 /d/project/opra-main-service/opra-main-service (master)
$ git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   controllers/riskissue/risk_issue_controller.go
        modified:   routes/riskissue.go
        modified:   services/riskissue/risk_issue_service.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        note.md