package main

const BasePath = "http://localhost:8080"
const CommonEndpointFormat = "%s%s"
const CommonQueryFormat = "%s?%s"
const SepaEndpoint = "/transactions/sepa"
const SwiftEndpoint = "/transactions/swift"
const PayerParam = "payer"
const SenderParam = "sender"
const ReceiverParam = "receiver"
const BeneficiaryParam = "beneficiary"
const AfterTimestamParam = "after-timestamp"
const ContentType = "Content-Type"
const JsonContentType = "application/json"
const CompanyIban = "company-iban"
const CompanyIDParam = "company-id"
const NumberOfNecessaryCallsForBalance = 4
