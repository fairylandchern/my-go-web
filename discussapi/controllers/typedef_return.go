package controllers

type STATUS int

const SUCCESS  STATUS=200

//ERROR
const DATAEXIST STATUS = 40001
const INSERTERR STATUS = 40002
const UPDATEERR STATUS = 40003
const DELETEERR STATUS = 40004
const DATANOTEXIST STATUS = 40005
const GETALLDATAERROR STATUS = 40006


