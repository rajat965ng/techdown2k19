provider "xray" {
  host = "http://35.194.21.117"
  port = "8000"
  username = "admin"
  password = "password"
}

resource "xray_blockdownload" "blockdown_one" {
  unscanned = true
  active = true
}

resource "xray_action" "actionone" {
  mails = ["abc@xyz.com"]
  fail_build = false
  block_download = "${xray_blockdownload.blockdown_one}"
}

resource "xray_criteria" "criteriaone" {
  "min_severity" =  "high"
}

resource "xray_rule" "rule_one" {
  name = "ruleone"
  priority = 1
  actions = "${xray_action.actionone}"
  criteria = "${xray_criteria.criteriaone}"
}


resource "xray_policy" "policy_one" {
  name = "securityPolicy"
  type = "security"
  description = "Yet another security policy"
  rules = ["${xray_rule.rule_one}"]
}