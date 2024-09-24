package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterProcessingStateDetail string

const (
	PrinterProcessingStateDetail_AlertRemovalOfBinaryChangeEntry           PrinterProcessingStateDetail = "alertRemovalOfBinaryChangeEntry"
	PrinterProcessingStateDetail_BanderAdded                               PrinterProcessingStateDetail = "banderAdded"
	PrinterProcessingStateDetail_BanderAlmostEmpty                         PrinterProcessingStateDetail = "banderAlmostEmpty"
	PrinterProcessingStateDetail_BanderAlmostFull                          PrinterProcessingStateDetail = "banderAlmostFull"
	PrinterProcessingStateDetail_BanderAtLimit                             PrinterProcessingStateDetail = "banderAtLimit"
	PrinterProcessingStateDetail_BanderClosed                              PrinterProcessingStateDetail = "banderClosed"
	PrinterProcessingStateDetail_BanderConfigurationChange                 PrinterProcessingStateDetail = "banderConfigurationChange"
	PrinterProcessingStateDetail_BanderCoverClosed                         PrinterProcessingStateDetail = "banderCoverClosed"
	PrinterProcessingStateDetail_BanderCoverOpen                           PrinterProcessingStateDetail = "banderCoverOpen"
	PrinterProcessingStateDetail_BanderEmpty                               PrinterProcessingStateDetail = "banderEmpty"
	PrinterProcessingStateDetail_BanderFull                                PrinterProcessingStateDetail = "banderFull"
	PrinterProcessingStateDetail_BanderInterlockClosed                     PrinterProcessingStateDetail = "banderInterlockClosed"
	PrinterProcessingStateDetail_BanderInterlockOpen                       PrinterProcessingStateDetail = "banderInterlockOpen"
	PrinterProcessingStateDetail_BanderJam                                 PrinterProcessingStateDetail = "banderJam"
	PrinterProcessingStateDetail_BanderLifeAlmostOver                      PrinterProcessingStateDetail = "banderLifeAlmostOver"
	PrinterProcessingStateDetail_BanderLifeOver                            PrinterProcessingStateDetail = "banderLifeOver"
	PrinterProcessingStateDetail_BanderMemoryExhausted                     PrinterProcessingStateDetail = "banderMemoryExhausted"
	PrinterProcessingStateDetail_BanderMissing                             PrinterProcessingStateDetail = "banderMissing"
	PrinterProcessingStateDetail_BanderMotorFailure                        PrinterProcessingStateDetail = "banderMotorFailure"
	PrinterProcessingStateDetail_BanderNearLimit                           PrinterProcessingStateDetail = "banderNearLimit"
	PrinterProcessingStateDetail_BanderOffline                             PrinterProcessingStateDetail = "banderOffline"
	PrinterProcessingStateDetail_BanderOpened                              PrinterProcessingStateDetail = "banderOpened"
	PrinterProcessingStateDetail_BanderOverTemperature                     PrinterProcessingStateDetail = "banderOverTemperature"
	PrinterProcessingStateDetail_BanderPowerSaver                          PrinterProcessingStateDetail = "banderPowerSaver"
	PrinterProcessingStateDetail_BanderRecoverableFailure                  PrinterProcessingStateDetail = "banderRecoverableFailure"
	PrinterProcessingStateDetail_BanderRecoverableStorage                  PrinterProcessingStateDetail = "banderRecoverableStorage"
	PrinterProcessingStateDetail_BanderRemoved                             PrinterProcessingStateDetail = "banderRemoved"
	PrinterProcessingStateDetail_BanderResourceAdded                       PrinterProcessingStateDetail = "banderResourceAdded"
	PrinterProcessingStateDetail_BanderResourceRemoved                     PrinterProcessingStateDetail = "banderResourceRemoved"
	PrinterProcessingStateDetail_BanderThermistorFailure                   PrinterProcessingStateDetail = "banderThermistorFailure"
	PrinterProcessingStateDetail_BanderTimingFailure                       PrinterProcessingStateDetail = "banderTimingFailure"
	PrinterProcessingStateDetail_BanderTurnedOff                           PrinterProcessingStateDetail = "banderTurnedOff"
	PrinterProcessingStateDetail_BanderTurnedOn                            PrinterProcessingStateDetail = "banderTurnedOn"
	PrinterProcessingStateDetail_BanderUnderTemperature                    PrinterProcessingStateDetail = "banderUnderTemperature"
	PrinterProcessingStateDetail_BanderUnrecoverableFailure                PrinterProcessingStateDetail = "banderUnrecoverableFailure"
	PrinterProcessingStateDetail_BanderUnrecoverableStorageError           PrinterProcessingStateDetail = "banderUnrecoverableStorageError"
	PrinterProcessingStateDetail_BanderWarmingUp                           PrinterProcessingStateDetail = "banderWarmingUp"
	PrinterProcessingStateDetail_BinderAdded                               PrinterProcessingStateDetail = "binderAdded"
	PrinterProcessingStateDetail_BinderAlmostEmpty                         PrinterProcessingStateDetail = "binderAlmostEmpty"
	PrinterProcessingStateDetail_BinderAlmostFull                          PrinterProcessingStateDetail = "binderAlmostFull"
	PrinterProcessingStateDetail_BinderAtLimit                             PrinterProcessingStateDetail = "binderAtLimit"
	PrinterProcessingStateDetail_BinderClosed                              PrinterProcessingStateDetail = "binderClosed"
	PrinterProcessingStateDetail_BinderConfigurationChange                 PrinterProcessingStateDetail = "binderConfigurationChange"
	PrinterProcessingStateDetail_BinderCoverClosed                         PrinterProcessingStateDetail = "binderCoverClosed"
	PrinterProcessingStateDetail_BinderCoverOpen                           PrinterProcessingStateDetail = "binderCoverOpen"
	PrinterProcessingStateDetail_BinderEmpty                               PrinterProcessingStateDetail = "binderEmpty"
	PrinterProcessingStateDetail_BinderFull                                PrinterProcessingStateDetail = "binderFull"
	PrinterProcessingStateDetail_BinderInterlockClosed                     PrinterProcessingStateDetail = "binderInterlockClosed"
	PrinterProcessingStateDetail_BinderInterlockOpen                       PrinterProcessingStateDetail = "binderInterlockOpen"
	PrinterProcessingStateDetail_BinderJam                                 PrinterProcessingStateDetail = "binderJam"
	PrinterProcessingStateDetail_BinderLifeAlmostOver                      PrinterProcessingStateDetail = "binderLifeAlmostOver"
	PrinterProcessingStateDetail_BinderLifeOver                            PrinterProcessingStateDetail = "binderLifeOver"
	PrinterProcessingStateDetail_BinderMemoryExhausted                     PrinterProcessingStateDetail = "binderMemoryExhausted"
	PrinterProcessingStateDetail_BinderMissing                             PrinterProcessingStateDetail = "binderMissing"
	PrinterProcessingStateDetail_BinderMotorFailure                        PrinterProcessingStateDetail = "binderMotorFailure"
	PrinterProcessingStateDetail_BinderNearLimit                           PrinterProcessingStateDetail = "binderNearLimit"
	PrinterProcessingStateDetail_BinderOffline                             PrinterProcessingStateDetail = "binderOffline"
	PrinterProcessingStateDetail_BinderOpened                              PrinterProcessingStateDetail = "binderOpened"
	PrinterProcessingStateDetail_BinderOverTemperature                     PrinterProcessingStateDetail = "binderOverTemperature"
	PrinterProcessingStateDetail_BinderPowerSaver                          PrinterProcessingStateDetail = "binderPowerSaver"
	PrinterProcessingStateDetail_BinderRecoverableFailure                  PrinterProcessingStateDetail = "binderRecoverableFailure"
	PrinterProcessingStateDetail_BinderRecoverableStorage                  PrinterProcessingStateDetail = "binderRecoverableStorage"
	PrinterProcessingStateDetail_BinderRemoved                             PrinterProcessingStateDetail = "binderRemoved"
	PrinterProcessingStateDetail_BinderResourceAdded                       PrinterProcessingStateDetail = "binderResourceAdded"
	PrinterProcessingStateDetail_BinderResourceRemoved                     PrinterProcessingStateDetail = "binderResourceRemoved"
	PrinterProcessingStateDetail_BinderThermistorFailure                   PrinterProcessingStateDetail = "binderThermistorFailure"
	PrinterProcessingStateDetail_BinderTimingFailure                       PrinterProcessingStateDetail = "binderTimingFailure"
	PrinterProcessingStateDetail_BinderTurnedOff                           PrinterProcessingStateDetail = "binderTurnedOff"
	PrinterProcessingStateDetail_BinderTurnedOn                            PrinterProcessingStateDetail = "binderTurnedOn"
	PrinterProcessingStateDetail_BinderUnderTemperature                    PrinterProcessingStateDetail = "binderUnderTemperature"
	PrinterProcessingStateDetail_BinderUnrecoverableFailure                PrinterProcessingStateDetail = "binderUnrecoverableFailure"
	PrinterProcessingStateDetail_BinderUnrecoverableStorageError           PrinterProcessingStateDetail = "binderUnrecoverableStorageError"
	PrinterProcessingStateDetail_BinderWarmingUp                           PrinterProcessingStateDetail = "binderWarmingUp"
	PrinterProcessingStateDetail_CameraFailure                             PrinterProcessingStateDetail = "cameraFailure"
	PrinterProcessingStateDetail_ChamberCooling                            PrinterProcessingStateDetail = "chamberCooling"
	PrinterProcessingStateDetail_ChamberFailure                            PrinterProcessingStateDetail = "chamberFailure"
	PrinterProcessingStateDetail_ChamberHeating                            PrinterProcessingStateDetail = "chamberHeating"
	PrinterProcessingStateDetail_ChamberTemperatureHigh                    PrinterProcessingStateDetail = "chamberTemperatureHigh"
	PrinterProcessingStateDetail_ChamberTemperatureLow                     PrinterProcessingStateDetail = "chamberTemperatureLow"
	PrinterProcessingStateDetail_CleanerLifeAlmostOver                     PrinterProcessingStateDetail = "cleanerLifeAlmostOver"
	PrinterProcessingStateDetail_CleanerLifeOver                           PrinterProcessingStateDetail = "cleanerLifeOver"
	PrinterProcessingStateDetail_ConfigurationChange                       PrinterProcessingStateDetail = "configurationChange"
	PrinterProcessingStateDetail_ConnectingToDevice                        PrinterProcessingStateDetail = "connectingToDevice"
	PrinterProcessingStateDetail_CoverOpen                                 PrinterProcessingStateDetail = "coverOpen"
	PrinterProcessingStateDetail_Deactivated                               PrinterProcessingStateDetail = "deactivated"
	PrinterProcessingStateDetail_Deleted                                   PrinterProcessingStateDetail = "deleted"
	PrinterProcessingStateDetail_DeveloperEmpty                            PrinterProcessingStateDetail = "developerEmpty"
	PrinterProcessingStateDetail_DeveloperLow                              PrinterProcessingStateDetail = "developerLow"
	PrinterProcessingStateDetail_DieCutterAdded                            PrinterProcessingStateDetail = "dieCutterAdded"
	PrinterProcessingStateDetail_DieCutterAlmostEmpty                      PrinterProcessingStateDetail = "dieCutterAlmostEmpty"
	PrinterProcessingStateDetail_DieCutterAlmostFull                       PrinterProcessingStateDetail = "dieCutterAlmostFull"
	PrinterProcessingStateDetail_DieCutterAtLimit                          PrinterProcessingStateDetail = "dieCutterAtLimit"
	PrinterProcessingStateDetail_DieCutterClosed                           PrinterProcessingStateDetail = "dieCutterClosed"
	PrinterProcessingStateDetail_DieCutterConfigurationChange              PrinterProcessingStateDetail = "dieCutterConfigurationChange"
	PrinterProcessingStateDetail_DieCutterCoverClosed                      PrinterProcessingStateDetail = "dieCutterCoverClosed"
	PrinterProcessingStateDetail_DieCutterCoverOpen                        PrinterProcessingStateDetail = "dieCutterCoverOpen"
	PrinterProcessingStateDetail_DieCutterEmpty                            PrinterProcessingStateDetail = "dieCutterEmpty"
	PrinterProcessingStateDetail_DieCutterFull                             PrinterProcessingStateDetail = "dieCutterFull"
	PrinterProcessingStateDetail_DieCutterInterlockClosed                  PrinterProcessingStateDetail = "dieCutterInterlockClosed"
	PrinterProcessingStateDetail_DieCutterInterlockOpen                    PrinterProcessingStateDetail = "dieCutterInterlockOpen"
	PrinterProcessingStateDetail_DieCutterJam                              PrinterProcessingStateDetail = "dieCutterJam"
	PrinterProcessingStateDetail_DieCutterLifeAlmostOver                   PrinterProcessingStateDetail = "dieCutterLifeAlmostOver"
	PrinterProcessingStateDetail_DieCutterLifeOver                         PrinterProcessingStateDetail = "dieCutterLifeOver"
	PrinterProcessingStateDetail_DieCutterMemoryExhausted                  PrinterProcessingStateDetail = "dieCutterMemoryExhausted"
	PrinterProcessingStateDetail_DieCutterMissing                          PrinterProcessingStateDetail = "dieCutterMissing"
	PrinterProcessingStateDetail_DieCutterMotorFailure                     PrinterProcessingStateDetail = "dieCutterMotorFailure"
	PrinterProcessingStateDetail_DieCutterNearLimit                        PrinterProcessingStateDetail = "dieCutterNearLimit"
	PrinterProcessingStateDetail_DieCutterOffline                          PrinterProcessingStateDetail = "dieCutterOffline"
	PrinterProcessingStateDetail_DieCutterOpened                           PrinterProcessingStateDetail = "dieCutterOpened"
	PrinterProcessingStateDetail_DieCutterOverTemperature                  PrinterProcessingStateDetail = "dieCutterOverTemperature"
	PrinterProcessingStateDetail_DieCutterPowerSaver                       PrinterProcessingStateDetail = "dieCutterPowerSaver"
	PrinterProcessingStateDetail_DieCutterRecoverableFailure               PrinterProcessingStateDetail = "dieCutterRecoverableFailure"
	PrinterProcessingStateDetail_DieCutterRecoverableStorage               PrinterProcessingStateDetail = "dieCutterRecoverableStorage"
	PrinterProcessingStateDetail_DieCutterRemoved                          PrinterProcessingStateDetail = "dieCutterRemoved"
	PrinterProcessingStateDetail_DieCutterResourceAdded                    PrinterProcessingStateDetail = "dieCutterResourceAdded"
	PrinterProcessingStateDetail_DieCutterResourceRemoved                  PrinterProcessingStateDetail = "dieCutterResourceRemoved"
	PrinterProcessingStateDetail_DieCutterThermistorFailure                PrinterProcessingStateDetail = "dieCutterThermistorFailure"
	PrinterProcessingStateDetail_DieCutterTimingFailure                    PrinterProcessingStateDetail = "dieCutterTimingFailure"
	PrinterProcessingStateDetail_DieCutterTurnedOff                        PrinterProcessingStateDetail = "dieCutterTurnedOff"
	PrinterProcessingStateDetail_DieCutterTurnedOn                         PrinterProcessingStateDetail = "dieCutterTurnedOn"
	PrinterProcessingStateDetail_DieCutterUnderTemperature                 PrinterProcessingStateDetail = "dieCutterUnderTemperature"
	PrinterProcessingStateDetail_DieCutterUnrecoverableFailure             PrinterProcessingStateDetail = "dieCutterUnrecoverableFailure"
	PrinterProcessingStateDetail_DieCutterUnrecoverableStorageError        PrinterProcessingStateDetail = "dieCutterUnrecoverableStorageError"
	PrinterProcessingStateDetail_DieCutterWarmingUp                        PrinterProcessingStateDetail = "dieCutterWarmingUp"
	PrinterProcessingStateDetail_DoorOpen                                  PrinterProcessingStateDetail = "doorOpen"
	PrinterProcessingStateDetail_ExtruderCooling                           PrinterProcessingStateDetail = "extruderCooling"
	PrinterProcessingStateDetail_ExtruderFailure                           PrinterProcessingStateDetail = "extruderFailure"
	PrinterProcessingStateDetail_ExtruderHeating                           PrinterProcessingStateDetail = "extruderHeating"
	PrinterProcessingStateDetail_ExtruderJam                               PrinterProcessingStateDetail = "extruderJam"
	PrinterProcessingStateDetail_ExtruderTemperatureHigh                   PrinterProcessingStateDetail = "extruderTemperatureHigh"
	PrinterProcessingStateDetail_ExtruderTemperatureLow                    PrinterProcessingStateDetail = "extruderTemperatureLow"
	PrinterProcessingStateDetail_FanFailure                                PrinterProcessingStateDetail = "fanFailure"
	PrinterProcessingStateDetail_FaxModemLifeAlmostOver                    PrinterProcessingStateDetail = "faxModemLifeAlmostOver"
	PrinterProcessingStateDetail_FaxModemLifeOver                          PrinterProcessingStateDetail = "faxModemLifeOver"
	PrinterProcessingStateDetail_FaxModemMissing                           PrinterProcessingStateDetail = "faxModemMissing"
	PrinterProcessingStateDetail_FaxModemTurnedOff                         PrinterProcessingStateDetail = "faxModemTurnedOff"
	PrinterProcessingStateDetail_FaxModemTurnedOn                          PrinterProcessingStateDetail = "faxModemTurnedOn"
	PrinterProcessingStateDetail_FolderAdded                               PrinterProcessingStateDetail = "folderAdded"
	PrinterProcessingStateDetail_FolderAlmostEmpty                         PrinterProcessingStateDetail = "folderAlmostEmpty"
	PrinterProcessingStateDetail_FolderAlmostFull                          PrinterProcessingStateDetail = "folderAlmostFull"
	PrinterProcessingStateDetail_FolderAtLimit                             PrinterProcessingStateDetail = "folderAtLimit"
	PrinterProcessingStateDetail_FolderClosed                              PrinterProcessingStateDetail = "folderClosed"
	PrinterProcessingStateDetail_FolderConfigurationChange                 PrinterProcessingStateDetail = "folderConfigurationChange"
	PrinterProcessingStateDetail_FolderCoverClosed                         PrinterProcessingStateDetail = "folderCoverClosed"
	PrinterProcessingStateDetail_FolderCoverOpen                           PrinterProcessingStateDetail = "folderCoverOpen"
	PrinterProcessingStateDetail_FolderEmpty                               PrinterProcessingStateDetail = "folderEmpty"
	PrinterProcessingStateDetail_FolderFull                                PrinterProcessingStateDetail = "folderFull"
	PrinterProcessingStateDetail_FolderInterlockClosed                     PrinterProcessingStateDetail = "folderInterlockClosed"
	PrinterProcessingStateDetail_FolderInterlockOpen                       PrinterProcessingStateDetail = "folderInterlockOpen"
	PrinterProcessingStateDetail_FolderJam                                 PrinterProcessingStateDetail = "folderJam"
	PrinterProcessingStateDetail_FolderLifeAlmostOver                      PrinterProcessingStateDetail = "folderLifeAlmostOver"
	PrinterProcessingStateDetail_FolderLifeOver                            PrinterProcessingStateDetail = "folderLifeOver"
	PrinterProcessingStateDetail_FolderMemoryExhausted                     PrinterProcessingStateDetail = "folderMemoryExhausted"
	PrinterProcessingStateDetail_FolderMissing                             PrinterProcessingStateDetail = "folderMissing"
	PrinterProcessingStateDetail_FolderMotorFailure                        PrinterProcessingStateDetail = "folderMotorFailure"
	PrinterProcessingStateDetail_FolderNearLimit                           PrinterProcessingStateDetail = "folderNearLimit"
	PrinterProcessingStateDetail_FolderOffline                             PrinterProcessingStateDetail = "folderOffline"
	PrinterProcessingStateDetail_FolderOpened                              PrinterProcessingStateDetail = "folderOpened"
	PrinterProcessingStateDetail_FolderOverTemperature                     PrinterProcessingStateDetail = "folderOverTemperature"
	PrinterProcessingStateDetail_FolderPowerSaver                          PrinterProcessingStateDetail = "folderPowerSaver"
	PrinterProcessingStateDetail_FolderRecoverableFailure                  PrinterProcessingStateDetail = "folderRecoverableFailure"
	PrinterProcessingStateDetail_FolderRecoverableStorage                  PrinterProcessingStateDetail = "folderRecoverableStorage"
	PrinterProcessingStateDetail_FolderRemoved                             PrinterProcessingStateDetail = "folderRemoved"
	PrinterProcessingStateDetail_FolderResourceAdded                       PrinterProcessingStateDetail = "folderResourceAdded"
	PrinterProcessingStateDetail_FolderResourceRemoved                     PrinterProcessingStateDetail = "folderResourceRemoved"
	PrinterProcessingStateDetail_FolderThermistorFailure                   PrinterProcessingStateDetail = "folderThermistorFailure"
	PrinterProcessingStateDetail_FolderTimingFailure                       PrinterProcessingStateDetail = "folderTimingFailure"
	PrinterProcessingStateDetail_FolderTurnedOff                           PrinterProcessingStateDetail = "folderTurnedOff"
	PrinterProcessingStateDetail_FolderTurnedOn                            PrinterProcessingStateDetail = "folderTurnedOn"
	PrinterProcessingStateDetail_FolderUnderTemperature                    PrinterProcessingStateDetail = "folderUnderTemperature"
	PrinterProcessingStateDetail_FolderUnrecoverableFailure                PrinterProcessingStateDetail = "folderUnrecoverableFailure"
	PrinterProcessingStateDetail_FolderUnrecoverableStorageError           PrinterProcessingStateDetail = "folderUnrecoverableStorageError"
	PrinterProcessingStateDetail_FolderWarmingUp                           PrinterProcessingStateDetail = "folderWarmingUp"
	PrinterProcessingStateDetail_FuserOverTemp                             PrinterProcessingStateDetail = "fuserOverTemp"
	PrinterProcessingStateDetail_FuserUnderTemp                            PrinterProcessingStateDetail = "fuserUnderTemp"
	PrinterProcessingStateDetail_Hibernate                                 PrinterProcessingStateDetail = "hibernate"
	PrinterProcessingStateDetail_HoldNewJobs                               PrinterProcessingStateDetail = "holdNewJobs"
	PrinterProcessingStateDetail_IdentifyPrinterRequested                  PrinterProcessingStateDetail = "identifyPrinterRequested"
	PrinterProcessingStateDetail_ImprinterAdded                            PrinterProcessingStateDetail = "imprinterAdded"
	PrinterProcessingStateDetail_ImprinterAlmostEmpty                      PrinterProcessingStateDetail = "imprinterAlmostEmpty"
	PrinterProcessingStateDetail_ImprinterAlmostFull                       PrinterProcessingStateDetail = "imprinterAlmostFull"
	PrinterProcessingStateDetail_ImprinterAtLimit                          PrinterProcessingStateDetail = "imprinterAtLimit"
	PrinterProcessingStateDetail_ImprinterClosed                           PrinterProcessingStateDetail = "imprinterClosed"
	PrinterProcessingStateDetail_ImprinterConfigurationChange              PrinterProcessingStateDetail = "imprinterConfigurationChange"
	PrinterProcessingStateDetail_ImprinterCoverClosed                      PrinterProcessingStateDetail = "imprinterCoverClosed"
	PrinterProcessingStateDetail_ImprinterCoverOpen                        PrinterProcessingStateDetail = "imprinterCoverOpen"
	PrinterProcessingStateDetail_ImprinterEmpty                            PrinterProcessingStateDetail = "imprinterEmpty"
	PrinterProcessingStateDetail_ImprinterFull                             PrinterProcessingStateDetail = "imprinterFull"
	PrinterProcessingStateDetail_ImprinterInterlockClosed                  PrinterProcessingStateDetail = "imprinterInterlockClosed"
	PrinterProcessingStateDetail_ImprinterInterlockOpen                    PrinterProcessingStateDetail = "imprinterInterlockOpen"
	PrinterProcessingStateDetail_ImprinterJam                              PrinterProcessingStateDetail = "imprinterJam"
	PrinterProcessingStateDetail_ImprinterLifeAlmostOver                   PrinterProcessingStateDetail = "imprinterLifeAlmostOver"
	PrinterProcessingStateDetail_ImprinterLifeOver                         PrinterProcessingStateDetail = "imprinterLifeOver"
	PrinterProcessingStateDetail_ImprinterMemoryExhausted                  PrinterProcessingStateDetail = "imprinterMemoryExhausted"
	PrinterProcessingStateDetail_ImprinterMissing                          PrinterProcessingStateDetail = "imprinterMissing"
	PrinterProcessingStateDetail_ImprinterMotorFailure                     PrinterProcessingStateDetail = "imprinterMotorFailure"
	PrinterProcessingStateDetail_ImprinterNearLimit                        PrinterProcessingStateDetail = "imprinterNearLimit"
	PrinterProcessingStateDetail_ImprinterOffline                          PrinterProcessingStateDetail = "imprinterOffline"
	PrinterProcessingStateDetail_ImprinterOpened                           PrinterProcessingStateDetail = "imprinterOpened"
	PrinterProcessingStateDetail_ImprinterOverTemperature                  PrinterProcessingStateDetail = "imprinterOverTemperature"
	PrinterProcessingStateDetail_ImprinterPowerSaver                       PrinterProcessingStateDetail = "imprinterPowerSaver"
	PrinterProcessingStateDetail_ImprinterRecoverableFailure               PrinterProcessingStateDetail = "imprinterRecoverableFailure"
	PrinterProcessingStateDetail_ImprinterRecoverableStorage               PrinterProcessingStateDetail = "imprinterRecoverableStorage"
	PrinterProcessingStateDetail_ImprinterRemoved                          PrinterProcessingStateDetail = "imprinterRemoved"
	PrinterProcessingStateDetail_ImprinterResourceAdded                    PrinterProcessingStateDetail = "imprinterResourceAdded"
	PrinterProcessingStateDetail_ImprinterResourceRemoved                  PrinterProcessingStateDetail = "imprinterResourceRemoved"
	PrinterProcessingStateDetail_ImprinterThermistorFailure                PrinterProcessingStateDetail = "imprinterThermistorFailure"
	PrinterProcessingStateDetail_ImprinterTimingFailure                    PrinterProcessingStateDetail = "imprinterTimingFailure"
	PrinterProcessingStateDetail_ImprinterTurnedOff                        PrinterProcessingStateDetail = "imprinterTurnedOff"
	PrinterProcessingStateDetail_ImprinterTurnedOn                         PrinterProcessingStateDetail = "imprinterTurnedOn"
	PrinterProcessingStateDetail_ImprinterUnderTemperature                 PrinterProcessingStateDetail = "imprinterUnderTemperature"
	PrinterProcessingStateDetail_ImprinterUnrecoverableFailure             PrinterProcessingStateDetail = "imprinterUnrecoverableFailure"
	PrinterProcessingStateDetail_ImprinterUnrecoverableStorageError        PrinterProcessingStateDetail = "imprinterUnrecoverableStorageError"
	PrinterProcessingStateDetail_ImprinterWarmingUp                        PrinterProcessingStateDetail = "imprinterWarmingUp"
	PrinterProcessingStateDetail_InputCannotFeedSizeSelected               PrinterProcessingStateDetail = "inputCannotFeedSizeSelected"
	PrinterProcessingStateDetail_InputManualInputRequest                   PrinterProcessingStateDetail = "inputManualInputRequest"
	PrinterProcessingStateDetail_InputMediaColorChange                     PrinterProcessingStateDetail = "inputMediaColorChange"
	PrinterProcessingStateDetail_InputMediaFormPartsChange                 PrinterProcessingStateDetail = "inputMediaFormPartsChange"
	PrinterProcessingStateDetail_InputMediaSizeChange                      PrinterProcessingStateDetail = "inputMediaSizeChange"
	PrinterProcessingStateDetail_InputMediaTrayFailure                     PrinterProcessingStateDetail = "inputMediaTrayFailure"
	PrinterProcessingStateDetail_InputMediaTrayFeedError                   PrinterProcessingStateDetail = "inputMediaTrayFeedError"
	PrinterProcessingStateDetail_InputMediaTrayJam                         PrinterProcessingStateDetail = "inputMediaTrayJam"
	PrinterProcessingStateDetail_InputMediaTypeChange                      PrinterProcessingStateDetail = "inputMediaTypeChange"
	PrinterProcessingStateDetail_InputMediaWeightChange                    PrinterProcessingStateDetail = "inputMediaWeightChange"
	PrinterProcessingStateDetail_InputPickRollerFailure                    PrinterProcessingStateDetail = "inputPickRollerFailure"
	PrinterProcessingStateDetail_InputPickRollerLifeOver                   PrinterProcessingStateDetail = "inputPickRollerLifeOver"
	PrinterProcessingStateDetail_InputPickRollerLifeWarn                   PrinterProcessingStateDetail = "inputPickRollerLifeWarn"
	PrinterProcessingStateDetail_InputPickRollerMissing                    PrinterProcessingStateDetail = "inputPickRollerMissing"
	PrinterProcessingStateDetail_InputTrayElevationFailure                 PrinterProcessingStateDetail = "inputTrayElevationFailure"
	PrinterProcessingStateDetail_InputTrayMissing                          PrinterProcessingStateDetail = "inputTrayMissing"
	PrinterProcessingStateDetail_InputTrayPositionFailure                  PrinterProcessingStateDetail = "inputTrayPositionFailure"
	PrinterProcessingStateDetail_InserterAdded                             PrinterProcessingStateDetail = "inserterAdded"
	PrinterProcessingStateDetail_InserterAlmostEmpty                       PrinterProcessingStateDetail = "inserterAlmostEmpty"
	PrinterProcessingStateDetail_InserterAlmostFull                        PrinterProcessingStateDetail = "inserterAlmostFull"
	PrinterProcessingStateDetail_InserterAtLimit                           PrinterProcessingStateDetail = "inserterAtLimit"
	PrinterProcessingStateDetail_InserterClosed                            PrinterProcessingStateDetail = "inserterClosed"
	PrinterProcessingStateDetail_InserterConfigurationChange               PrinterProcessingStateDetail = "inserterConfigurationChange"
	PrinterProcessingStateDetail_InserterCoverClosed                       PrinterProcessingStateDetail = "inserterCoverClosed"
	PrinterProcessingStateDetail_InserterCoverOpen                         PrinterProcessingStateDetail = "inserterCoverOpen"
	PrinterProcessingStateDetail_InserterEmpty                             PrinterProcessingStateDetail = "inserterEmpty"
	PrinterProcessingStateDetail_InserterFull                              PrinterProcessingStateDetail = "inserterFull"
	PrinterProcessingStateDetail_InserterInterlockClosed                   PrinterProcessingStateDetail = "inserterInterlockClosed"
	PrinterProcessingStateDetail_InserterInterlockOpen                     PrinterProcessingStateDetail = "inserterInterlockOpen"
	PrinterProcessingStateDetail_InserterJam                               PrinterProcessingStateDetail = "inserterJam"
	PrinterProcessingStateDetail_InserterLifeAlmostOver                    PrinterProcessingStateDetail = "inserterLifeAlmostOver"
	PrinterProcessingStateDetail_InserterLifeOver                          PrinterProcessingStateDetail = "inserterLifeOver"
	PrinterProcessingStateDetail_InserterMemoryExhausted                   PrinterProcessingStateDetail = "inserterMemoryExhausted"
	PrinterProcessingStateDetail_InserterMissing                           PrinterProcessingStateDetail = "inserterMissing"
	PrinterProcessingStateDetail_InserterMotorFailure                      PrinterProcessingStateDetail = "inserterMotorFailure"
	PrinterProcessingStateDetail_InserterNearLimit                         PrinterProcessingStateDetail = "inserterNearLimit"
	PrinterProcessingStateDetail_InserterOffline                           PrinterProcessingStateDetail = "inserterOffline"
	PrinterProcessingStateDetail_InserterOpened                            PrinterProcessingStateDetail = "inserterOpened"
	PrinterProcessingStateDetail_InserterOverTemperature                   PrinterProcessingStateDetail = "inserterOverTemperature"
	PrinterProcessingStateDetail_InserterPowerSaver                        PrinterProcessingStateDetail = "inserterPowerSaver"
	PrinterProcessingStateDetail_InserterRecoverableFailure                PrinterProcessingStateDetail = "inserterRecoverableFailure"
	PrinterProcessingStateDetail_InserterRecoverableStorage                PrinterProcessingStateDetail = "inserterRecoverableStorage"
	PrinterProcessingStateDetail_InserterRemoved                           PrinterProcessingStateDetail = "inserterRemoved"
	PrinterProcessingStateDetail_InserterResourceAdded                     PrinterProcessingStateDetail = "inserterResourceAdded"
	PrinterProcessingStateDetail_InserterResourceRemoved                   PrinterProcessingStateDetail = "inserterResourceRemoved"
	PrinterProcessingStateDetail_InserterThermistorFailure                 PrinterProcessingStateDetail = "inserterThermistorFailure"
	PrinterProcessingStateDetail_InserterTimingFailure                     PrinterProcessingStateDetail = "inserterTimingFailure"
	PrinterProcessingStateDetail_InserterTurnedOff                         PrinterProcessingStateDetail = "inserterTurnedOff"
	PrinterProcessingStateDetail_InserterTurnedOn                          PrinterProcessingStateDetail = "inserterTurnedOn"
	PrinterProcessingStateDetail_InserterUnderTemperature                  PrinterProcessingStateDetail = "inserterUnderTemperature"
	PrinterProcessingStateDetail_InserterUnrecoverableFailure              PrinterProcessingStateDetail = "inserterUnrecoverableFailure"
	PrinterProcessingStateDetail_InserterUnrecoverableStorageError         PrinterProcessingStateDetail = "inserterUnrecoverableStorageError"
	PrinterProcessingStateDetail_InserterWarmingUp                         PrinterProcessingStateDetail = "inserterWarmingUp"
	PrinterProcessingStateDetail_InterlockClosed                           PrinterProcessingStateDetail = "interlockClosed"
	PrinterProcessingStateDetail_InterlockOpen                             PrinterProcessingStateDetail = "interlockOpen"
	PrinterProcessingStateDetail_InterpreterCartridgeAdded                 PrinterProcessingStateDetail = "interpreterCartridgeAdded"
	PrinterProcessingStateDetail_InterpreterCartridgeDeleted               PrinterProcessingStateDetail = "interpreterCartridgeDeleted"
	PrinterProcessingStateDetail_InterpreterComplexPageEncountered         PrinterProcessingStateDetail = "interpreterComplexPageEncountered"
	PrinterProcessingStateDetail_InterpreterMemoryDecrease                 PrinterProcessingStateDetail = "interpreterMemoryDecrease"
	PrinterProcessingStateDetail_InterpreterMemoryIncrease                 PrinterProcessingStateDetail = "interpreterMemoryIncrease"
	PrinterProcessingStateDetail_InterpreterResourceAdded                  PrinterProcessingStateDetail = "interpreterResourceAdded"
	PrinterProcessingStateDetail_InterpreterResourceDeleted                PrinterProcessingStateDetail = "interpreterResourceDeleted"
	PrinterProcessingStateDetail_InterpreterResourceUnavailable            PrinterProcessingStateDetail = "interpreterResourceUnavailable"
	PrinterProcessingStateDetail_LampAtEol                                 PrinterProcessingStateDetail = "lampAtEol"
	PrinterProcessingStateDetail_LampFailure                               PrinterProcessingStateDetail = "lampFailure"
	PrinterProcessingStateDetail_LampNearEol                               PrinterProcessingStateDetail = "lampNearEol"
	PrinterProcessingStateDetail_LaserAtEol                                PrinterProcessingStateDetail = "laserAtEol"
	PrinterProcessingStateDetail_LaserFailure                              PrinterProcessingStateDetail = "laserFailure"
	PrinterProcessingStateDetail_LaserNearEol                              PrinterProcessingStateDetail = "laserNearEol"
	PrinterProcessingStateDetail_MakeEnvelopeAdded                         PrinterProcessingStateDetail = "makeEnvelopeAdded"
	PrinterProcessingStateDetail_MakeEnvelopeAlmostEmpty                   PrinterProcessingStateDetail = "makeEnvelopeAlmostEmpty"
	PrinterProcessingStateDetail_MakeEnvelopeAlmostFull                    PrinterProcessingStateDetail = "makeEnvelopeAlmostFull"
	PrinterProcessingStateDetail_MakeEnvelopeAtLimit                       PrinterProcessingStateDetail = "makeEnvelopeAtLimit"
	PrinterProcessingStateDetail_MakeEnvelopeClosed                        PrinterProcessingStateDetail = "makeEnvelopeClosed"
	PrinterProcessingStateDetail_MakeEnvelopeConfigurationChange           PrinterProcessingStateDetail = "makeEnvelopeConfigurationChange"
	PrinterProcessingStateDetail_MakeEnvelopeCoverClosed                   PrinterProcessingStateDetail = "makeEnvelopeCoverClosed"
	PrinterProcessingStateDetail_MakeEnvelopeCoverOpen                     PrinterProcessingStateDetail = "makeEnvelopeCoverOpen"
	PrinterProcessingStateDetail_MakeEnvelopeEmpty                         PrinterProcessingStateDetail = "makeEnvelopeEmpty"
	PrinterProcessingStateDetail_MakeEnvelopeFull                          PrinterProcessingStateDetail = "makeEnvelopeFull"
	PrinterProcessingStateDetail_MakeEnvelopeInterlockClosed               PrinterProcessingStateDetail = "makeEnvelopeInterlockClosed"
	PrinterProcessingStateDetail_MakeEnvelopeInterlockOpen                 PrinterProcessingStateDetail = "makeEnvelopeInterlockOpen"
	PrinterProcessingStateDetail_MakeEnvelopeJam                           PrinterProcessingStateDetail = "makeEnvelopeJam"
	PrinterProcessingStateDetail_MakeEnvelopeLifeAlmostOver                PrinterProcessingStateDetail = "makeEnvelopeLifeAlmostOver"
	PrinterProcessingStateDetail_MakeEnvelopeLifeOver                      PrinterProcessingStateDetail = "makeEnvelopeLifeOver"
	PrinterProcessingStateDetail_MakeEnvelopeMemoryExhausted               PrinterProcessingStateDetail = "makeEnvelopeMemoryExhausted"
	PrinterProcessingStateDetail_MakeEnvelopeMissing                       PrinterProcessingStateDetail = "makeEnvelopeMissing"
	PrinterProcessingStateDetail_MakeEnvelopeMotorFailure                  PrinterProcessingStateDetail = "makeEnvelopeMotorFailure"
	PrinterProcessingStateDetail_MakeEnvelopeNearLimit                     PrinterProcessingStateDetail = "makeEnvelopeNearLimit"
	PrinterProcessingStateDetail_MakeEnvelopeOffline                       PrinterProcessingStateDetail = "makeEnvelopeOffline"
	PrinterProcessingStateDetail_MakeEnvelopeOpened                        PrinterProcessingStateDetail = "makeEnvelopeOpened"
	PrinterProcessingStateDetail_MakeEnvelopeOverTemperature               PrinterProcessingStateDetail = "makeEnvelopeOverTemperature"
	PrinterProcessingStateDetail_MakeEnvelopePowerSaver                    PrinterProcessingStateDetail = "makeEnvelopePowerSaver"
	PrinterProcessingStateDetail_MakeEnvelopeRecoverableFailure            PrinterProcessingStateDetail = "makeEnvelopeRecoverableFailure"
	PrinterProcessingStateDetail_MakeEnvelopeRecoverableStorage            PrinterProcessingStateDetail = "makeEnvelopeRecoverableStorage"
	PrinterProcessingStateDetail_MakeEnvelopeRemoved                       PrinterProcessingStateDetail = "makeEnvelopeRemoved"
	PrinterProcessingStateDetail_MakeEnvelopeResourceAdded                 PrinterProcessingStateDetail = "makeEnvelopeResourceAdded"
	PrinterProcessingStateDetail_MakeEnvelopeResourceRemoved               PrinterProcessingStateDetail = "makeEnvelopeResourceRemoved"
	PrinterProcessingStateDetail_MakeEnvelopeThermistorFailure             PrinterProcessingStateDetail = "makeEnvelopeThermistorFailure"
	PrinterProcessingStateDetail_MakeEnvelopeTimingFailure                 PrinterProcessingStateDetail = "makeEnvelopeTimingFailure"
	PrinterProcessingStateDetail_MakeEnvelopeTurnedOff                     PrinterProcessingStateDetail = "makeEnvelopeTurnedOff"
	PrinterProcessingStateDetail_MakeEnvelopeTurnedOn                      PrinterProcessingStateDetail = "makeEnvelopeTurnedOn"
	PrinterProcessingStateDetail_MakeEnvelopeUnderTemperature              PrinterProcessingStateDetail = "makeEnvelopeUnderTemperature"
	PrinterProcessingStateDetail_MakeEnvelopeUnrecoverableFailure          PrinterProcessingStateDetail = "makeEnvelopeUnrecoverableFailure"
	PrinterProcessingStateDetail_MakeEnvelopeUnrecoverableStorageError     PrinterProcessingStateDetail = "makeEnvelopeUnrecoverableStorageError"
	PrinterProcessingStateDetail_MakeEnvelopeWarmingUp                     PrinterProcessingStateDetail = "makeEnvelopeWarmingUp"
	PrinterProcessingStateDetail_MarkerAdjustingPrintQuality               PrinterProcessingStateDetail = "markerAdjustingPrintQuality"
	PrinterProcessingStateDetail_MarkerCleanerMissing                      PrinterProcessingStateDetail = "markerCleanerMissing"
	PrinterProcessingStateDetail_MarkerDeveloperAlmostEmpty                PrinterProcessingStateDetail = "markerDeveloperAlmostEmpty"
	PrinterProcessingStateDetail_MarkerDeveloperEmpty                      PrinterProcessingStateDetail = "markerDeveloperEmpty"
	PrinterProcessingStateDetail_MarkerDeveloperMissing                    PrinterProcessingStateDetail = "markerDeveloperMissing"
	PrinterProcessingStateDetail_MarkerFuserMissing                        PrinterProcessingStateDetail = "markerFuserMissing"
	PrinterProcessingStateDetail_MarkerFuserThermistorFailure              PrinterProcessingStateDetail = "markerFuserThermistorFailure"
	PrinterProcessingStateDetail_MarkerFuserTimingFailure                  PrinterProcessingStateDetail = "markerFuserTimingFailure"
	PrinterProcessingStateDetail_MarkerInkAlmostEmpty                      PrinterProcessingStateDetail = "markerInkAlmostEmpty"
	PrinterProcessingStateDetail_MarkerInkEmpty                            PrinterProcessingStateDetail = "markerInkEmpty"
	PrinterProcessingStateDetail_MarkerInkMissing                          PrinterProcessingStateDetail = "markerInkMissing"
	PrinterProcessingStateDetail_MarkerOpcMissing                          PrinterProcessingStateDetail = "markerOpcMissing"
	PrinterProcessingStateDetail_MarkerPrintRibbonAlmostEmpty              PrinterProcessingStateDetail = "markerPrintRibbonAlmostEmpty"
	PrinterProcessingStateDetail_MarkerPrintRibbonEmpty                    PrinterProcessingStateDetail = "markerPrintRibbonEmpty"
	PrinterProcessingStateDetail_MarkerPrintRibbonMissing                  PrinterProcessingStateDetail = "markerPrintRibbonMissing"
	PrinterProcessingStateDetail_MarkerSupplyAlmostEmpty                   PrinterProcessingStateDetail = "markerSupplyAlmostEmpty"
	PrinterProcessingStateDetail_MarkerSupplyEmpty                         PrinterProcessingStateDetail = "markerSupplyEmpty"
	PrinterProcessingStateDetail_MarkerSupplyLow                           PrinterProcessingStateDetail = "markerSupplyLow"
	PrinterProcessingStateDetail_MarkerSupplyMissing                       PrinterProcessingStateDetail = "markerSupplyMissing"
	PrinterProcessingStateDetail_MarkerTonerCartridgeMissing               PrinterProcessingStateDetail = "markerTonerCartridgeMissing"
	PrinterProcessingStateDetail_MarkerTonerMissing                        PrinterProcessingStateDetail = "markerTonerMissing"
	PrinterProcessingStateDetail_MarkerWasteAlmostFull                     PrinterProcessingStateDetail = "markerWasteAlmostFull"
	PrinterProcessingStateDetail_MarkerWasteFull                           PrinterProcessingStateDetail = "markerWasteFull"
	PrinterProcessingStateDetail_MarkerWasteInkReceptacleAlmostFull        PrinterProcessingStateDetail = "markerWasteInkReceptacleAlmostFull"
	PrinterProcessingStateDetail_MarkerWasteInkReceptacleFull              PrinterProcessingStateDetail = "markerWasteInkReceptacleFull"
	PrinterProcessingStateDetail_MarkerWasteInkReceptacleMissing           PrinterProcessingStateDetail = "markerWasteInkReceptacleMissing"
	PrinterProcessingStateDetail_MarkerWasteMissing                        PrinterProcessingStateDetail = "markerWasteMissing"
	PrinterProcessingStateDetail_MarkerWasteTonerReceptacleAlmostFull      PrinterProcessingStateDetail = "markerWasteTonerReceptacleAlmostFull"
	PrinterProcessingStateDetail_MarkerWasteTonerReceptacleFull            PrinterProcessingStateDetail = "markerWasteTonerReceptacleFull"
	PrinterProcessingStateDetail_MarkerWasteTonerReceptacleMissing         PrinterProcessingStateDetail = "markerWasteTonerReceptacleMissing"
	PrinterProcessingStateDetail_MaterialEmpty                             PrinterProcessingStateDetail = "materialEmpty"
	PrinterProcessingStateDetail_MaterialLow                               PrinterProcessingStateDetail = "materialLow"
	PrinterProcessingStateDetail_MaterialNeeded                            PrinterProcessingStateDetail = "materialNeeded"
	PrinterProcessingStateDetail_MediaDrying                               PrinterProcessingStateDetail = "mediaDrying"
	PrinterProcessingStateDetail_MediaEmpty                                PrinterProcessingStateDetail = "mediaEmpty"
	PrinterProcessingStateDetail_MediaJam                                  PrinterProcessingStateDetail = "mediaJam"
	PrinterProcessingStateDetail_MediaLow                                  PrinterProcessingStateDetail = "mediaLow"
	PrinterProcessingStateDetail_MediaNeeded                               PrinterProcessingStateDetail = "mediaNeeded"
	PrinterProcessingStateDetail_MediaPathCannotDuplexMediaSelected        PrinterProcessingStateDetail = "mediaPathCannotDuplexMediaSelected"
	PrinterProcessingStateDetail_MediaPathFailure                          PrinterProcessingStateDetail = "mediaPathFailure"
	PrinterProcessingStateDetail_MediaPathInputEmpty                       PrinterProcessingStateDetail = "mediaPathInputEmpty"
	PrinterProcessingStateDetail_MediaPathInputFeedError                   PrinterProcessingStateDetail = "mediaPathInputFeedError"
	PrinterProcessingStateDetail_MediaPathInputJam                         PrinterProcessingStateDetail = "mediaPathInputJam"
	PrinterProcessingStateDetail_MediaPathInputRequest                     PrinterProcessingStateDetail = "mediaPathInputRequest"
	PrinterProcessingStateDetail_MediaPathJam                              PrinterProcessingStateDetail = "mediaPathJam"
	PrinterProcessingStateDetail_MediaPathMediaTrayAlmostFull              PrinterProcessingStateDetail = "mediaPathMediaTrayAlmostFull"
	PrinterProcessingStateDetail_MediaPathMediaTrayFull                    PrinterProcessingStateDetail = "mediaPathMediaTrayFull"
	PrinterProcessingStateDetail_MediaPathMediaTrayMissing                 PrinterProcessingStateDetail = "mediaPathMediaTrayMissing"
	PrinterProcessingStateDetail_MediaPathOutputFeedError                  PrinterProcessingStateDetail = "mediaPathOutputFeedError"
	PrinterProcessingStateDetail_MediaPathOutputFull                       PrinterProcessingStateDetail = "mediaPathOutputFull"
	PrinterProcessingStateDetail_MediaPathOutputJam                        PrinterProcessingStateDetail = "mediaPathOutputJam"
	PrinterProcessingStateDetail_MediaPathPickRollerFailure                PrinterProcessingStateDetail = "mediaPathPickRollerFailure"
	PrinterProcessingStateDetail_MediaPathPickRollerLifeOver               PrinterProcessingStateDetail = "mediaPathPickRollerLifeOver"
	PrinterProcessingStateDetail_MediaPathPickRollerLifeWarn               PrinterProcessingStateDetail = "mediaPathPickRollerLifeWarn"
	PrinterProcessingStateDetail_MediaPathPickRollerMissing                PrinterProcessingStateDetail = "mediaPathPickRollerMissing"
	PrinterProcessingStateDetail_MotorFailure                              PrinterProcessingStateDetail = "motorFailure"
	PrinterProcessingStateDetail_MovingToPaused                            PrinterProcessingStateDetail = "movingToPaused"
	PrinterProcessingStateDetail_None                                      PrinterProcessingStateDetail = "none"
	PrinterProcessingStateDetail_OpticalPhotoConductorLifeOver             PrinterProcessingStateDetail = "opticalPhotoConductorLifeOver"
	PrinterProcessingStateDetail_OpticalPhotoConductorNearEndOfLife        PrinterProcessingStateDetail = "opticalPhotoConductorNearEndOfLife"
	PrinterProcessingStateDetail_Other                                     PrinterProcessingStateDetail = "other"
	PrinterProcessingStateDetail_OutputAreaAlmostFull                      PrinterProcessingStateDetail = "outputAreaAlmostFull"
	PrinterProcessingStateDetail_OutputAreaFull                            PrinterProcessingStateDetail = "outputAreaFull"
	PrinterProcessingStateDetail_OutputMailboxSelectFailure                PrinterProcessingStateDetail = "outputMailboxSelectFailure"
	PrinterProcessingStateDetail_OutputMediaTrayFailure                    PrinterProcessingStateDetail = "outputMediaTrayFailure"
	PrinterProcessingStateDetail_OutputMediaTrayFeedError                  PrinterProcessingStateDetail = "outputMediaTrayFeedError"
	PrinterProcessingStateDetail_OutputMediaTrayJam                        PrinterProcessingStateDetail = "outputMediaTrayJam"
	PrinterProcessingStateDetail_OutputTrayMissing                         PrinterProcessingStateDetail = "outputTrayMissing"
	PrinterProcessingStateDetail_Paused                                    PrinterProcessingStateDetail = "paused"
	PrinterProcessingStateDetail_PerforaterAdded                           PrinterProcessingStateDetail = "perforaterAdded"
	PrinterProcessingStateDetail_PerforaterAlmostEmpty                     PrinterProcessingStateDetail = "perforaterAlmostEmpty"
	PrinterProcessingStateDetail_PerforaterAlmostFull                      PrinterProcessingStateDetail = "perforaterAlmostFull"
	PrinterProcessingStateDetail_PerforaterAtLimit                         PrinterProcessingStateDetail = "perforaterAtLimit"
	PrinterProcessingStateDetail_PerforaterClosed                          PrinterProcessingStateDetail = "perforaterClosed"
	PrinterProcessingStateDetail_PerforaterConfigurationChange             PrinterProcessingStateDetail = "perforaterConfigurationChange"
	PrinterProcessingStateDetail_PerforaterCoverClosed                     PrinterProcessingStateDetail = "perforaterCoverClosed"
	PrinterProcessingStateDetail_PerforaterCoverOpen                       PrinterProcessingStateDetail = "perforaterCoverOpen"
	PrinterProcessingStateDetail_PerforaterEmpty                           PrinterProcessingStateDetail = "perforaterEmpty"
	PrinterProcessingStateDetail_PerforaterFull                            PrinterProcessingStateDetail = "perforaterFull"
	PrinterProcessingStateDetail_PerforaterInterlockClosed                 PrinterProcessingStateDetail = "perforaterInterlockClosed"
	PrinterProcessingStateDetail_PerforaterInterlockOpen                   PrinterProcessingStateDetail = "perforaterInterlockOpen"
	PrinterProcessingStateDetail_PerforaterJam                             PrinterProcessingStateDetail = "perforaterJam"
	PrinterProcessingStateDetail_PerforaterLifeAlmostOver                  PrinterProcessingStateDetail = "perforaterLifeAlmostOver"
	PrinterProcessingStateDetail_PerforaterLifeOver                        PrinterProcessingStateDetail = "perforaterLifeOver"
	PrinterProcessingStateDetail_PerforaterMemoryExhausted                 PrinterProcessingStateDetail = "perforaterMemoryExhausted"
	PrinterProcessingStateDetail_PerforaterMissing                         PrinterProcessingStateDetail = "perforaterMissing"
	PrinterProcessingStateDetail_PerforaterMotorFailure                    PrinterProcessingStateDetail = "perforaterMotorFailure"
	PrinterProcessingStateDetail_PerforaterNearLimit                       PrinterProcessingStateDetail = "perforaterNearLimit"
	PrinterProcessingStateDetail_PerforaterOffline                         PrinterProcessingStateDetail = "perforaterOffline"
	PrinterProcessingStateDetail_PerforaterOpened                          PrinterProcessingStateDetail = "perforaterOpened"
	PrinterProcessingStateDetail_PerforaterOverTemperature                 PrinterProcessingStateDetail = "perforaterOverTemperature"
	PrinterProcessingStateDetail_PerforaterPowerSaver                      PrinterProcessingStateDetail = "perforaterPowerSaver"
	PrinterProcessingStateDetail_PerforaterRecoverableFailure              PrinterProcessingStateDetail = "perforaterRecoverableFailure"
	PrinterProcessingStateDetail_PerforaterRecoverableStorage              PrinterProcessingStateDetail = "perforaterRecoverableStorage"
	PrinterProcessingStateDetail_PerforaterRemoved                         PrinterProcessingStateDetail = "perforaterRemoved"
	PrinterProcessingStateDetail_PerforaterResourceAdded                   PrinterProcessingStateDetail = "perforaterResourceAdded"
	PrinterProcessingStateDetail_PerforaterResourceRemoved                 PrinterProcessingStateDetail = "perforaterResourceRemoved"
	PrinterProcessingStateDetail_PerforaterThermistorFailure               PrinterProcessingStateDetail = "perforaterThermistorFailure"
	PrinterProcessingStateDetail_PerforaterTimingFailure                   PrinterProcessingStateDetail = "perforaterTimingFailure"
	PrinterProcessingStateDetail_PerforaterTurnedOff                       PrinterProcessingStateDetail = "perforaterTurnedOff"
	PrinterProcessingStateDetail_PerforaterTurnedOn                        PrinterProcessingStateDetail = "perforaterTurnedOn"
	PrinterProcessingStateDetail_PerforaterUnderTemperature                PrinterProcessingStateDetail = "perforaterUnderTemperature"
	PrinterProcessingStateDetail_PerforaterUnrecoverableFailure            PrinterProcessingStateDetail = "perforaterUnrecoverableFailure"
	PrinterProcessingStateDetail_PerforaterUnrecoverableStorageError       PrinterProcessingStateDetail = "perforaterUnrecoverableStorageError"
	PrinterProcessingStateDetail_PerforaterWarmingUp                       PrinterProcessingStateDetail = "perforaterWarmingUp"
	PrinterProcessingStateDetail_PlatformCooling                           PrinterProcessingStateDetail = "platformCooling"
	PrinterProcessingStateDetail_PlatformFailure                           PrinterProcessingStateDetail = "platformFailure"
	PrinterProcessingStateDetail_PlatformHeating                           PrinterProcessingStateDetail = "platformHeating"
	PrinterProcessingStateDetail_PlatformTemperatureHigh                   PrinterProcessingStateDetail = "platformTemperatureHigh"
	PrinterProcessingStateDetail_PlatformTemperatureLow                    PrinterProcessingStateDetail = "platformTemperatureLow"
	PrinterProcessingStateDetail_PowerDown                                 PrinterProcessingStateDetail = "powerDown"
	PrinterProcessingStateDetail_PowerUp                                   PrinterProcessingStateDetail = "powerUp"
	PrinterProcessingStateDetail_PrinterManualReset                        PrinterProcessingStateDetail = "printerManualReset"
	PrinterProcessingStateDetail_PrinterNmsReset                           PrinterProcessingStateDetail = "printerNmsReset"
	PrinterProcessingStateDetail_PrinterReadyToPrint                       PrinterProcessingStateDetail = "printerReadyToPrint"
	PrinterProcessingStateDetail_PuncherAdded                              PrinterProcessingStateDetail = "puncherAdded"
	PrinterProcessingStateDetail_PuncherAlmostEmpty                        PrinterProcessingStateDetail = "puncherAlmostEmpty"
	PrinterProcessingStateDetail_PuncherAlmostFull                         PrinterProcessingStateDetail = "puncherAlmostFull"
	PrinterProcessingStateDetail_PuncherAtLimit                            PrinterProcessingStateDetail = "puncherAtLimit"
	PrinterProcessingStateDetail_PuncherClosed                             PrinterProcessingStateDetail = "puncherClosed"
	PrinterProcessingStateDetail_PuncherConfigurationChange                PrinterProcessingStateDetail = "puncherConfigurationChange"
	PrinterProcessingStateDetail_PuncherCoverClosed                        PrinterProcessingStateDetail = "puncherCoverClosed"
	PrinterProcessingStateDetail_PuncherCoverOpen                          PrinterProcessingStateDetail = "puncherCoverOpen"
	PrinterProcessingStateDetail_PuncherEmpty                              PrinterProcessingStateDetail = "puncherEmpty"
	PrinterProcessingStateDetail_PuncherFull                               PrinterProcessingStateDetail = "puncherFull"
	PrinterProcessingStateDetail_PuncherInterlockClosed                    PrinterProcessingStateDetail = "puncherInterlockClosed"
	PrinterProcessingStateDetail_PuncherInterlockOpen                      PrinterProcessingStateDetail = "puncherInterlockOpen"
	PrinterProcessingStateDetail_PuncherJam                                PrinterProcessingStateDetail = "puncherJam"
	PrinterProcessingStateDetail_PuncherLifeAlmostOver                     PrinterProcessingStateDetail = "puncherLifeAlmostOver"
	PrinterProcessingStateDetail_PuncherLifeOver                           PrinterProcessingStateDetail = "puncherLifeOver"
	PrinterProcessingStateDetail_PuncherMemoryExhausted                    PrinterProcessingStateDetail = "puncherMemoryExhausted"
	PrinterProcessingStateDetail_PuncherMissing                            PrinterProcessingStateDetail = "puncherMissing"
	PrinterProcessingStateDetail_PuncherMotorFailure                       PrinterProcessingStateDetail = "puncherMotorFailure"
	PrinterProcessingStateDetail_PuncherNearLimit                          PrinterProcessingStateDetail = "puncherNearLimit"
	PrinterProcessingStateDetail_PuncherOffline                            PrinterProcessingStateDetail = "puncherOffline"
	PrinterProcessingStateDetail_PuncherOpened                             PrinterProcessingStateDetail = "puncherOpened"
	PrinterProcessingStateDetail_PuncherOverTemperature                    PrinterProcessingStateDetail = "puncherOverTemperature"
	PrinterProcessingStateDetail_PuncherPowerSaver                         PrinterProcessingStateDetail = "puncherPowerSaver"
	PrinterProcessingStateDetail_PuncherRecoverableFailure                 PrinterProcessingStateDetail = "puncherRecoverableFailure"
	PrinterProcessingStateDetail_PuncherRecoverableStorage                 PrinterProcessingStateDetail = "puncherRecoverableStorage"
	PrinterProcessingStateDetail_PuncherRemoved                            PrinterProcessingStateDetail = "puncherRemoved"
	PrinterProcessingStateDetail_PuncherResourceAdded                      PrinterProcessingStateDetail = "puncherResourceAdded"
	PrinterProcessingStateDetail_PuncherResourceRemoved                    PrinterProcessingStateDetail = "puncherResourceRemoved"
	PrinterProcessingStateDetail_PuncherThermistorFailure                  PrinterProcessingStateDetail = "puncherThermistorFailure"
	PrinterProcessingStateDetail_PuncherTimingFailure                      PrinterProcessingStateDetail = "puncherTimingFailure"
	PrinterProcessingStateDetail_PuncherTurnedOff                          PrinterProcessingStateDetail = "puncherTurnedOff"
	PrinterProcessingStateDetail_PuncherTurnedOn                           PrinterProcessingStateDetail = "puncherTurnedOn"
	PrinterProcessingStateDetail_PuncherUnderTemperature                   PrinterProcessingStateDetail = "puncherUnderTemperature"
	PrinterProcessingStateDetail_PuncherUnrecoverableFailure               PrinterProcessingStateDetail = "puncherUnrecoverableFailure"
	PrinterProcessingStateDetail_PuncherUnrecoverableStorageError          PrinterProcessingStateDetail = "puncherUnrecoverableStorageError"
	PrinterProcessingStateDetail_PuncherWarmingUp                          PrinterProcessingStateDetail = "puncherWarmingUp"
	PrinterProcessingStateDetail_Resuming                                  PrinterProcessingStateDetail = "resuming"
	PrinterProcessingStateDetail_ScanMediaPathFailure                      PrinterProcessingStateDetail = "scanMediaPathFailure"
	PrinterProcessingStateDetail_ScanMediaPathInputEmpty                   PrinterProcessingStateDetail = "scanMediaPathInputEmpty"
	PrinterProcessingStateDetail_ScanMediaPathInputFeedError               PrinterProcessingStateDetail = "scanMediaPathInputFeedError"
	PrinterProcessingStateDetail_ScanMediaPathInputJam                     PrinterProcessingStateDetail = "scanMediaPathInputJam"
	PrinterProcessingStateDetail_ScanMediaPathInputRequest                 PrinterProcessingStateDetail = "scanMediaPathInputRequest"
	PrinterProcessingStateDetail_ScanMediaPathJam                          PrinterProcessingStateDetail = "scanMediaPathJam"
	PrinterProcessingStateDetail_ScanMediaPathOutputFeedError              PrinterProcessingStateDetail = "scanMediaPathOutputFeedError"
	PrinterProcessingStateDetail_ScanMediaPathOutputFull                   PrinterProcessingStateDetail = "scanMediaPathOutputFull"
	PrinterProcessingStateDetail_ScanMediaPathOutputJam                    PrinterProcessingStateDetail = "scanMediaPathOutputJam"
	PrinterProcessingStateDetail_ScanMediaPathPickRollerFailure            PrinterProcessingStateDetail = "scanMediaPathPickRollerFailure"
	PrinterProcessingStateDetail_ScanMediaPathPickRollerLifeOver           PrinterProcessingStateDetail = "scanMediaPathPickRollerLifeOver"
	PrinterProcessingStateDetail_ScanMediaPathPickRollerLifeWarn           PrinterProcessingStateDetail = "scanMediaPathPickRollerLifeWarn"
	PrinterProcessingStateDetail_ScanMediaPathPickRollerMissing            PrinterProcessingStateDetail = "scanMediaPathPickRollerMissing"
	PrinterProcessingStateDetail_ScanMediaPathTrayAlmostFull               PrinterProcessingStateDetail = "scanMediaPathTrayAlmostFull"
	PrinterProcessingStateDetail_ScanMediaPathTrayFull                     PrinterProcessingStateDetail = "scanMediaPathTrayFull"
	PrinterProcessingStateDetail_ScanMediaPathTrayMissing                  PrinterProcessingStateDetail = "scanMediaPathTrayMissing"
	PrinterProcessingStateDetail_ScannerLightFailure                       PrinterProcessingStateDetail = "scannerLightFailure"
	PrinterProcessingStateDetail_ScannerLightLifeAlmostOver                PrinterProcessingStateDetail = "scannerLightLifeAlmostOver"
	PrinterProcessingStateDetail_ScannerLightLifeOver                      PrinterProcessingStateDetail = "scannerLightLifeOver"
	PrinterProcessingStateDetail_ScannerLightMissing                       PrinterProcessingStateDetail = "scannerLightMissing"
	PrinterProcessingStateDetail_ScannerSensorFailure                      PrinterProcessingStateDetail = "scannerSensorFailure"
	PrinterProcessingStateDetail_ScannerSensorLifeAlmostOver               PrinterProcessingStateDetail = "scannerSensorLifeAlmostOver"
	PrinterProcessingStateDetail_ScannerSensorLifeOver                     PrinterProcessingStateDetail = "scannerSensorLifeOver"
	PrinterProcessingStateDetail_ScannerSensorMissing                      PrinterProcessingStateDetail = "scannerSensorMissing"
	PrinterProcessingStateDetail_SeparationCutterAdded                     PrinterProcessingStateDetail = "separationCutterAdded"
	PrinterProcessingStateDetail_SeparationCutterAlmostEmpty               PrinterProcessingStateDetail = "separationCutterAlmostEmpty"
	PrinterProcessingStateDetail_SeparationCutterAlmostFull                PrinterProcessingStateDetail = "separationCutterAlmostFull"
	PrinterProcessingStateDetail_SeparationCutterAtLimit                   PrinterProcessingStateDetail = "separationCutterAtLimit"
	PrinterProcessingStateDetail_SeparationCutterClosed                    PrinterProcessingStateDetail = "separationCutterClosed"
	PrinterProcessingStateDetail_SeparationCutterConfigurationChange       PrinterProcessingStateDetail = "separationCutterConfigurationChange"
	PrinterProcessingStateDetail_SeparationCutterCoverClosed               PrinterProcessingStateDetail = "separationCutterCoverClosed"
	PrinterProcessingStateDetail_SeparationCutterCoverOpen                 PrinterProcessingStateDetail = "separationCutterCoverOpen"
	PrinterProcessingStateDetail_SeparationCutterEmpty                     PrinterProcessingStateDetail = "separationCutterEmpty"
	PrinterProcessingStateDetail_SeparationCutterFull                      PrinterProcessingStateDetail = "separationCutterFull"
	PrinterProcessingStateDetail_SeparationCutterInterlockClosed           PrinterProcessingStateDetail = "separationCutterInterlockClosed"
	PrinterProcessingStateDetail_SeparationCutterInterlockOpen             PrinterProcessingStateDetail = "separationCutterInterlockOpen"
	PrinterProcessingStateDetail_SeparationCutterJam                       PrinterProcessingStateDetail = "separationCutterJam"
	PrinterProcessingStateDetail_SeparationCutterLifeAlmostOver            PrinterProcessingStateDetail = "separationCutterLifeAlmostOver"
	PrinterProcessingStateDetail_SeparationCutterLifeOver                  PrinterProcessingStateDetail = "separationCutterLifeOver"
	PrinterProcessingStateDetail_SeparationCutterMemoryExhausted           PrinterProcessingStateDetail = "separationCutterMemoryExhausted"
	PrinterProcessingStateDetail_SeparationCutterMissing                   PrinterProcessingStateDetail = "separationCutterMissing"
	PrinterProcessingStateDetail_SeparationCutterMotorFailure              PrinterProcessingStateDetail = "separationCutterMotorFailure"
	PrinterProcessingStateDetail_SeparationCutterNearLimit                 PrinterProcessingStateDetail = "separationCutterNearLimit"
	PrinterProcessingStateDetail_SeparationCutterOffline                   PrinterProcessingStateDetail = "separationCutterOffline"
	PrinterProcessingStateDetail_SeparationCutterOpened                    PrinterProcessingStateDetail = "separationCutterOpened"
	PrinterProcessingStateDetail_SeparationCutterOverTemperature           PrinterProcessingStateDetail = "separationCutterOverTemperature"
	PrinterProcessingStateDetail_SeparationCutterPowerSaver                PrinterProcessingStateDetail = "separationCutterPowerSaver"
	PrinterProcessingStateDetail_SeparationCutterRecoverableFailure        PrinterProcessingStateDetail = "separationCutterRecoverableFailure"
	PrinterProcessingStateDetail_SeparationCutterRecoverableStorage        PrinterProcessingStateDetail = "separationCutterRecoverableStorage"
	PrinterProcessingStateDetail_SeparationCutterRemoved                   PrinterProcessingStateDetail = "separationCutterRemoved"
	PrinterProcessingStateDetail_SeparationCutterResourceAdded             PrinterProcessingStateDetail = "separationCutterResourceAdded"
	PrinterProcessingStateDetail_SeparationCutterResourceRemoved           PrinterProcessingStateDetail = "separationCutterResourceRemoved"
	PrinterProcessingStateDetail_SeparationCutterThermistorFailure         PrinterProcessingStateDetail = "separationCutterThermistorFailure"
	PrinterProcessingStateDetail_SeparationCutterTimingFailure             PrinterProcessingStateDetail = "separationCutterTimingFailure"
	PrinterProcessingStateDetail_SeparationCutterTurnedOff                 PrinterProcessingStateDetail = "separationCutterTurnedOff"
	PrinterProcessingStateDetail_SeparationCutterTurnedOn                  PrinterProcessingStateDetail = "separationCutterTurnedOn"
	PrinterProcessingStateDetail_SeparationCutterUnderTemperature          PrinterProcessingStateDetail = "separationCutterUnderTemperature"
	PrinterProcessingStateDetail_SeparationCutterUnrecoverableFailure      PrinterProcessingStateDetail = "separationCutterUnrecoverableFailure"
	PrinterProcessingStateDetail_SeparationCutterUnrecoverableStorageError PrinterProcessingStateDetail = "separationCutterUnrecoverableStorageError"
	PrinterProcessingStateDetail_SeparationCutterWarmingUp                 PrinterProcessingStateDetail = "separationCutterWarmingUp"
	PrinterProcessingStateDetail_SheetRotatorAdded                         PrinterProcessingStateDetail = "sheetRotatorAdded"
	PrinterProcessingStateDetail_SheetRotatorAlmostEmpty                   PrinterProcessingStateDetail = "sheetRotatorAlmostEmpty"
	PrinterProcessingStateDetail_SheetRotatorAlmostFull                    PrinterProcessingStateDetail = "sheetRotatorAlmostFull"
	PrinterProcessingStateDetail_SheetRotatorAtLimit                       PrinterProcessingStateDetail = "sheetRotatorAtLimit"
	PrinterProcessingStateDetail_SheetRotatorClosed                        PrinterProcessingStateDetail = "sheetRotatorClosed"
	PrinterProcessingStateDetail_SheetRotatorConfigurationChange           PrinterProcessingStateDetail = "sheetRotatorConfigurationChange"
	PrinterProcessingStateDetail_SheetRotatorCoverClosed                   PrinterProcessingStateDetail = "sheetRotatorCoverClosed"
	PrinterProcessingStateDetail_SheetRotatorCoverOpen                     PrinterProcessingStateDetail = "sheetRotatorCoverOpen"
	PrinterProcessingStateDetail_SheetRotatorEmpty                         PrinterProcessingStateDetail = "sheetRotatorEmpty"
	PrinterProcessingStateDetail_SheetRotatorFull                          PrinterProcessingStateDetail = "sheetRotatorFull"
	PrinterProcessingStateDetail_SheetRotatorInterlockClosed               PrinterProcessingStateDetail = "sheetRotatorInterlockClosed"
	PrinterProcessingStateDetail_SheetRotatorInterlockOpen                 PrinterProcessingStateDetail = "sheetRotatorInterlockOpen"
	PrinterProcessingStateDetail_SheetRotatorJam                           PrinterProcessingStateDetail = "sheetRotatorJam"
	PrinterProcessingStateDetail_SheetRotatorLifeAlmostOver                PrinterProcessingStateDetail = "sheetRotatorLifeAlmostOver"
	PrinterProcessingStateDetail_SheetRotatorLifeOver                      PrinterProcessingStateDetail = "sheetRotatorLifeOver"
	PrinterProcessingStateDetail_SheetRotatorMemoryExhausted               PrinterProcessingStateDetail = "sheetRotatorMemoryExhausted"
	PrinterProcessingStateDetail_SheetRotatorMissing                       PrinterProcessingStateDetail = "sheetRotatorMissing"
	PrinterProcessingStateDetail_SheetRotatorMotorFailure                  PrinterProcessingStateDetail = "sheetRotatorMotorFailure"
	PrinterProcessingStateDetail_SheetRotatorNearLimit                     PrinterProcessingStateDetail = "sheetRotatorNearLimit"
	PrinterProcessingStateDetail_SheetRotatorOffline                       PrinterProcessingStateDetail = "sheetRotatorOffline"
	PrinterProcessingStateDetail_SheetRotatorOpened                        PrinterProcessingStateDetail = "sheetRotatorOpened"
	PrinterProcessingStateDetail_SheetRotatorOverTemperature               PrinterProcessingStateDetail = "sheetRotatorOverTemperature"
	PrinterProcessingStateDetail_SheetRotatorPowerSaver                    PrinterProcessingStateDetail = "sheetRotatorPowerSaver"
	PrinterProcessingStateDetail_SheetRotatorRecoverableFailure            PrinterProcessingStateDetail = "sheetRotatorRecoverableFailure"
	PrinterProcessingStateDetail_SheetRotatorRecoverableStorage            PrinterProcessingStateDetail = "sheetRotatorRecoverableStorage"
	PrinterProcessingStateDetail_SheetRotatorRemoved                       PrinterProcessingStateDetail = "sheetRotatorRemoved"
	PrinterProcessingStateDetail_SheetRotatorResourceAdded                 PrinterProcessingStateDetail = "sheetRotatorResourceAdded"
	PrinterProcessingStateDetail_SheetRotatorResourceRemoved               PrinterProcessingStateDetail = "sheetRotatorResourceRemoved"
	PrinterProcessingStateDetail_SheetRotatorThermistorFailure             PrinterProcessingStateDetail = "sheetRotatorThermistorFailure"
	PrinterProcessingStateDetail_SheetRotatorTimingFailure                 PrinterProcessingStateDetail = "sheetRotatorTimingFailure"
	PrinterProcessingStateDetail_SheetRotatorTurnedOff                     PrinterProcessingStateDetail = "sheetRotatorTurnedOff"
	PrinterProcessingStateDetail_SheetRotatorTurnedOn                      PrinterProcessingStateDetail = "sheetRotatorTurnedOn"
	PrinterProcessingStateDetail_SheetRotatorUnderTemperature              PrinterProcessingStateDetail = "sheetRotatorUnderTemperature"
	PrinterProcessingStateDetail_SheetRotatorUnrecoverableFailure          PrinterProcessingStateDetail = "sheetRotatorUnrecoverableFailure"
	PrinterProcessingStateDetail_SheetRotatorUnrecoverableStorageError     PrinterProcessingStateDetail = "sheetRotatorUnrecoverableStorageError"
	PrinterProcessingStateDetail_SheetRotatorWarmingUp                     PrinterProcessingStateDetail = "sheetRotatorWarmingUp"
	PrinterProcessingStateDetail_Shutdown                                  PrinterProcessingStateDetail = "shutdown"
	PrinterProcessingStateDetail_SlitterAdded                              PrinterProcessingStateDetail = "slitterAdded"
	PrinterProcessingStateDetail_SlitterAlmostEmpty                        PrinterProcessingStateDetail = "slitterAlmostEmpty"
	PrinterProcessingStateDetail_SlitterAlmostFull                         PrinterProcessingStateDetail = "slitterAlmostFull"
	PrinterProcessingStateDetail_SlitterAtLimit                            PrinterProcessingStateDetail = "slitterAtLimit"
	PrinterProcessingStateDetail_SlitterClosed                             PrinterProcessingStateDetail = "slitterClosed"
	PrinterProcessingStateDetail_SlitterConfigurationChange                PrinterProcessingStateDetail = "slitterConfigurationChange"
	PrinterProcessingStateDetail_SlitterCoverClosed                        PrinterProcessingStateDetail = "slitterCoverClosed"
	PrinterProcessingStateDetail_SlitterCoverOpen                          PrinterProcessingStateDetail = "slitterCoverOpen"
	PrinterProcessingStateDetail_SlitterEmpty                              PrinterProcessingStateDetail = "slitterEmpty"
	PrinterProcessingStateDetail_SlitterFull                               PrinterProcessingStateDetail = "slitterFull"
	PrinterProcessingStateDetail_SlitterInterlockClosed                    PrinterProcessingStateDetail = "slitterInterlockClosed"
	PrinterProcessingStateDetail_SlitterInterlockOpen                      PrinterProcessingStateDetail = "slitterInterlockOpen"
	PrinterProcessingStateDetail_SlitterJam                                PrinterProcessingStateDetail = "slitterJam"
	PrinterProcessingStateDetail_SlitterLifeAlmostOver                     PrinterProcessingStateDetail = "slitterLifeAlmostOver"
	PrinterProcessingStateDetail_SlitterLifeOver                           PrinterProcessingStateDetail = "slitterLifeOver"
	PrinterProcessingStateDetail_SlitterMemoryExhausted                    PrinterProcessingStateDetail = "slitterMemoryExhausted"
	PrinterProcessingStateDetail_SlitterMissing                            PrinterProcessingStateDetail = "slitterMissing"
	PrinterProcessingStateDetail_SlitterMotorFailure                       PrinterProcessingStateDetail = "slitterMotorFailure"
	PrinterProcessingStateDetail_SlitterNearLimit                          PrinterProcessingStateDetail = "slitterNearLimit"
	PrinterProcessingStateDetail_SlitterOffline                            PrinterProcessingStateDetail = "slitterOffline"
	PrinterProcessingStateDetail_SlitterOpened                             PrinterProcessingStateDetail = "slitterOpened"
	PrinterProcessingStateDetail_SlitterOverTemperature                    PrinterProcessingStateDetail = "slitterOverTemperature"
	PrinterProcessingStateDetail_SlitterPowerSaver                         PrinterProcessingStateDetail = "slitterPowerSaver"
	PrinterProcessingStateDetail_SlitterRecoverableFailure                 PrinterProcessingStateDetail = "slitterRecoverableFailure"
	PrinterProcessingStateDetail_SlitterRecoverableStorage                 PrinterProcessingStateDetail = "slitterRecoverableStorage"
	PrinterProcessingStateDetail_SlitterRemoved                            PrinterProcessingStateDetail = "slitterRemoved"
	PrinterProcessingStateDetail_SlitterResourceAdded                      PrinterProcessingStateDetail = "slitterResourceAdded"
	PrinterProcessingStateDetail_SlitterResourceRemoved                    PrinterProcessingStateDetail = "slitterResourceRemoved"
	PrinterProcessingStateDetail_SlitterThermistorFailure                  PrinterProcessingStateDetail = "slitterThermistorFailure"
	PrinterProcessingStateDetail_SlitterTimingFailure                      PrinterProcessingStateDetail = "slitterTimingFailure"
	PrinterProcessingStateDetail_SlitterTurnedOff                          PrinterProcessingStateDetail = "slitterTurnedOff"
	PrinterProcessingStateDetail_SlitterTurnedOn                           PrinterProcessingStateDetail = "slitterTurnedOn"
	PrinterProcessingStateDetail_SlitterUnderTemperature                   PrinterProcessingStateDetail = "slitterUnderTemperature"
	PrinterProcessingStateDetail_SlitterUnrecoverableFailure               PrinterProcessingStateDetail = "slitterUnrecoverableFailure"
	PrinterProcessingStateDetail_SlitterUnrecoverableStorageError          PrinterProcessingStateDetail = "slitterUnrecoverableStorageError"
	PrinterProcessingStateDetail_SlitterWarmingUp                          PrinterProcessingStateDetail = "slitterWarmingUp"
	PrinterProcessingStateDetail_SpoolAreaFull                             PrinterProcessingStateDetail = "spoolAreaFull"
	PrinterProcessingStateDetail_StackerAdded                              PrinterProcessingStateDetail = "stackerAdded"
	PrinterProcessingStateDetail_StackerAlmostEmpty                        PrinterProcessingStateDetail = "stackerAlmostEmpty"
	PrinterProcessingStateDetail_StackerAlmostFull                         PrinterProcessingStateDetail = "stackerAlmostFull"
	PrinterProcessingStateDetail_StackerAtLimit                            PrinterProcessingStateDetail = "stackerAtLimit"
	PrinterProcessingStateDetail_StackerClosed                             PrinterProcessingStateDetail = "stackerClosed"
	PrinterProcessingStateDetail_StackerConfigurationChange                PrinterProcessingStateDetail = "stackerConfigurationChange"
	PrinterProcessingStateDetail_StackerCoverClosed                        PrinterProcessingStateDetail = "stackerCoverClosed"
	PrinterProcessingStateDetail_StackerCoverOpen                          PrinterProcessingStateDetail = "stackerCoverOpen"
	PrinterProcessingStateDetail_StackerEmpty                              PrinterProcessingStateDetail = "stackerEmpty"
	PrinterProcessingStateDetail_StackerFull                               PrinterProcessingStateDetail = "stackerFull"
	PrinterProcessingStateDetail_StackerInterlockClosed                    PrinterProcessingStateDetail = "stackerInterlockClosed"
	PrinterProcessingStateDetail_StackerInterlockOpen                      PrinterProcessingStateDetail = "stackerInterlockOpen"
	PrinterProcessingStateDetail_StackerJam                                PrinterProcessingStateDetail = "stackerJam"
	PrinterProcessingStateDetail_StackerLifeAlmostOver                     PrinterProcessingStateDetail = "stackerLifeAlmostOver"
	PrinterProcessingStateDetail_StackerLifeOver                           PrinterProcessingStateDetail = "stackerLifeOver"
	PrinterProcessingStateDetail_StackerMemoryExhausted                    PrinterProcessingStateDetail = "stackerMemoryExhausted"
	PrinterProcessingStateDetail_StackerMissing                            PrinterProcessingStateDetail = "stackerMissing"
	PrinterProcessingStateDetail_StackerMotorFailure                       PrinterProcessingStateDetail = "stackerMotorFailure"
	PrinterProcessingStateDetail_StackerNearLimit                          PrinterProcessingStateDetail = "stackerNearLimit"
	PrinterProcessingStateDetail_StackerOffline                            PrinterProcessingStateDetail = "stackerOffline"
	PrinterProcessingStateDetail_StackerOpened                             PrinterProcessingStateDetail = "stackerOpened"
	PrinterProcessingStateDetail_StackerOverTemperature                    PrinterProcessingStateDetail = "stackerOverTemperature"
	PrinterProcessingStateDetail_StackerPowerSaver                         PrinterProcessingStateDetail = "stackerPowerSaver"
	PrinterProcessingStateDetail_StackerRecoverableFailure                 PrinterProcessingStateDetail = "stackerRecoverableFailure"
	PrinterProcessingStateDetail_StackerRecoverableStorage                 PrinterProcessingStateDetail = "stackerRecoverableStorage"
	PrinterProcessingStateDetail_StackerRemoved                            PrinterProcessingStateDetail = "stackerRemoved"
	PrinterProcessingStateDetail_StackerResourceAdded                      PrinterProcessingStateDetail = "stackerResourceAdded"
	PrinterProcessingStateDetail_StackerResourceRemoved                    PrinterProcessingStateDetail = "stackerResourceRemoved"
	PrinterProcessingStateDetail_StackerThermistorFailure                  PrinterProcessingStateDetail = "stackerThermistorFailure"
	PrinterProcessingStateDetail_StackerTimingFailure                      PrinterProcessingStateDetail = "stackerTimingFailure"
	PrinterProcessingStateDetail_StackerTurnedOff                          PrinterProcessingStateDetail = "stackerTurnedOff"
	PrinterProcessingStateDetail_StackerTurnedOn                           PrinterProcessingStateDetail = "stackerTurnedOn"
	PrinterProcessingStateDetail_StackerUnderTemperature                   PrinterProcessingStateDetail = "stackerUnderTemperature"
	PrinterProcessingStateDetail_StackerUnrecoverableFailure               PrinterProcessingStateDetail = "stackerUnrecoverableFailure"
	PrinterProcessingStateDetail_StackerUnrecoverableStorageError          PrinterProcessingStateDetail = "stackerUnrecoverableStorageError"
	PrinterProcessingStateDetail_StackerWarmingUp                          PrinterProcessingStateDetail = "stackerWarmingUp"
	PrinterProcessingStateDetail_Standby                                   PrinterProcessingStateDetail = "standby"
	PrinterProcessingStateDetail_StaplerAdded                              PrinterProcessingStateDetail = "staplerAdded"
	PrinterProcessingStateDetail_StaplerAlmostEmpty                        PrinterProcessingStateDetail = "staplerAlmostEmpty"
	PrinterProcessingStateDetail_StaplerAlmostFull                         PrinterProcessingStateDetail = "staplerAlmostFull"
	PrinterProcessingStateDetail_StaplerAtLimit                            PrinterProcessingStateDetail = "staplerAtLimit"
	PrinterProcessingStateDetail_StaplerClosed                             PrinterProcessingStateDetail = "staplerClosed"
	PrinterProcessingStateDetail_StaplerConfigurationChange                PrinterProcessingStateDetail = "staplerConfigurationChange"
	PrinterProcessingStateDetail_StaplerCoverClosed                        PrinterProcessingStateDetail = "staplerCoverClosed"
	PrinterProcessingStateDetail_StaplerCoverOpen                          PrinterProcessingStateDetail = "staplerCoverOpen"
	PrinterProcessingStateDetail_StaplerEmpty                              PrinterProcessingStateDetail = "staplerEmpty"
	PrinterProcessingStateDetail_StaplerFull                               PrinterProcessingStateDetail = "staplerFull"
	PrinterProcessingStateDetail_StaplerInterlockClosed                    PrinterProcessingStateDetail = "staplerInterlockClosed"
	PrinterProcessingStateDetail_StaplerInterlockOpen                      PrinterProcessingStateDetail = "staplerInterlockOpen"
	PrinterProcessingStateDetail_StaplerJam                                PrinterProcessingStateDetail = "staplerJam"
	PrinterProcessingStateDetail_StaplerLifeAlmostOver                     PrinterProcessingStateDetail = "staplerLifeAlmostOver"
	PrinterProcessingStateDetail_StaplerLifeOver                           PrinterProcessingStateDetail = "staplerLifeOver"
	PrinterProcessingStateDetail_StaplerMemoryExhausted                    PrinterProcessingStateDetail = "staplerMemoryExhausted"
	PrinterProcessingStateDetail_StaplerMissing                            PrinterProcessingStateDetail = "staplerMissing"
	PrinterProcessingStateDetail_StaplerMotorFailure                       PrinterProcessingStateDetail = "staplerMotorFailure"
	PrinterProcessingStateDetail_StaplerNearLimit                          PrinterProcessingStateDetail = "staplerNearLimit"
	PrinterProcessingStateDetail_StaplerOffline                            PrinterProcessingStateDetail = "staplerOffline"
	PrinterProcessingStateDetail_StaplerOpened                             PrinterProcessingStateDetail = "staplerOpened"
	PrinterProcessingStateDetail_StaplerOverTemperature                    PrinterProcessingStateDetail = "staplerOverTemperature"
	PrinterProcessingStateDetail_StaplerPowerSaver                         PrinterProcessingStateDetail = "staplerPowerSaver"
	PrinterProcessingStateDetail_StaplerRecoverableFailure                 PrinterProcessingStateDetail = "staplerRecoverableFailure"
	PrinterProcessingStateDetail_StaplerRecoverableStorage                 PrinterProcessingStateDetail = "staplerRecoverableStorage"
	PrinterProcessingStateDetail_StaplerRemoved                            PrinterProcessingStateDetail = "staplerRemoved"
	PrinterProcessingStateDetail_StaplerResourceAdded                      PrinterProcessingStateDetail = "staplerResourceAdded"
	PrinterProcessingStateDetail_StaplerResourceRemoved                    PrinterProcessingStateDetail = "staplerResourceRemoved"
	PrinterProcessingStateDetail_StaplerThermistorFailure                  PrinterProcessingStateDetail = "staplerThermistorFailure"
	PrinterProcessingStateDetail_StaplerTimingFailure                      PrinterProcessingStateDetail = "staplerTimingFailure"
	PrinterProcessingStateDetail_StaplerTurnedOff                          PrinterProcessingStateDetail = "staplerTurnedOff"
	PrinterProcessingStateDetail_StaplerTurnedOn                           PrinterProcessingStateDetail = "staplerTurnedOn"
	PrinterProcessingStateDetail_StaplerUnderTemperature                   PrinterProcessingStateDetail = "staplerUnderTemperature"
	PrinterProcessingStateDetail_StaplerUnrecoverableFailure               PrinterProcessingStateDetail = "staplerUnrecoverableFailure"
	PrinterProcessingStateDetail_StaplerUnrecoverableStorageError          PrinterProcessingStateDetail = "staplerUnrecoverableStorageError"
	PrinterProcessingStateDetail_StaplerWarmingUp                          PrinterProcessingStateDetail = "staplerWarmingUp"
	PrinterProcessingStateDetail_StitcherAdded                             PrinterProcessingStateDetail = "stitcherAdded"
	PrinterProcessingStateDetail_StitcherAlmostEmpty                       PrinterProcessingStateDetail = "stitcherAlmostEmpty"
	PrinterProcessingStateDetail_StitcherAlmostFull                        PrinterProcessingStateDetail = "stitcherAlmostFull"
	PrinterProcessingStateDetail_StitcherAtLimit                           PrinterProcessingStateDetail = "stitcherAtLimit"
	PrinterProcessingStateDetail_StitcherClosed                            PrinterProcessingStateDetail = "stitcherClosed"
	PrinterProcessingStateDetail_StitcherConfigurationChange               PrinterProcessingStateDetail = "stitcherConfigurationChange"
	PrinterProcessingStateDetail_StitcherCoverClosed                       PrinterProcessingStateDetail = "stitcherCoverClosed"
	PrinterProcessingStateDetail_StitcherCoverOpen                         PrinterProcessingStateDetail = "stitcherCoverOpen"
	PrinterProcessingStateDetail_StitcherEmpty                             PrinterProcessingStateDetail = "stitcherEmpty"
	PrinterProcessingStateDetail_StitcherFull                              PrinterProcessingStateDetail = "stitcherFull"
	PrinterProcessingStateDetail_StitcherInterlockClosed                   PrinterProcessingStateDetail = "stitcherInterlockClosed"
	PrinterProcessingStateDetail_StitcherInterlockOpen                     PrinterProcessingStateDetail = "stitcherInterlockOpen"
	PrinterProcessingStateDetail_StitcherJam                               PrinterProcessingStateDetail = "stitcherJam"
	PrinterProcessingStateDetail_StitcherLifeAlmostOver                    PrinterProcessingStateDetail = "stitcherLifeAlmostOver"
	PrinterProcessingStateDetail_StitcherLifeOver                          PrinterProcessingStateDetail = "stitcherLifeOver"
	PrinterProcessingStateDetail_StitcherMemoryExhausted                   PrinterProcessingStateDetail = "stitcherMemoryExhausted"
	PrinterProcessingStateDetail_StitcherMissing                           PrinterProcessingStateDetail = "stitcherMissing"
	PrinterProcessingStateDetail_StitcherMotorFailure                      PrinterProcessingStateDetail = "stitcherMotorFailure"
	PrinterProcessingStateDetail_StitcherNearLimit                         PrinterProcessingStateDetail = "stitcherNearLimit"
	PrinterProcessingStateDetail_StitcherOffline                           PrinterProcessingStateDetail = "stitcherOffline"
	PrinterProcessingStateDetail_StitcherOpened                            PrinterProcessingStateDetail = "stitcherOpened"
	PrinterProcessingStateDetail_StitcherOverTemperature                   PrinterProcessingStateDetail = "stitcherOverTemperature"
	PrinterProcessingStateDetail_StitcherPowerSaver                        PrinterProcessingStateDetail = "stitcherPowerSaver"
	PrinterProcessingStateDetail_StitcherRecoverableFailure                PrinterProcessingStateDetail = "stitcherRecoverableFailure"
	PrinterProcessingStateDetail_StitcherRecoverableStorage                PrinterProcessingStateDetail = "stitcherRecoverableStorage"
	PrinterProcessingStateDetail_StitcherRemoved                           PrinterProcessingStateDetail = "stitcherRemoved"
	PrinterProcessingStateDetail_StitcherResourceAdded                     PrinterProcessingStateDetail = "stitcherResourceAdded"
	PrinterProcessingStateDetail_StitcherResourceRemoved                   PrinterProcessingStateDetail = "stitcherResourceRemoved"
	PrinterProcessingStateDetail_StitcherThermistorFailure                 PrinterProcessingStateDetail = "stitcherThermistorFailure"
	PrinterProcessingStateDetail_StitcherTimingFailure                     PrinterProcessingStateDetail = "stitcherTimingFailure"
	PrinterProcessingStateDetail_StitcherTurnedOff                         PrinterProcessingStateDetail = "stitcherTurnedOff"
	PrinterProcessingStateDetail_StitcherTurnedOn                          PrinterProcessingStateDetail = "stitcherTurnedOn"
	PrinterProcessingStateDetail_StitcherUnderTemperature                  PrinterProcessingStateDetail = "stitcherUnderTemperature"
	PrinterProcessingStateDetail_StitcherUnrecoverableFailure              PrinterProcessingStateDetail = "stitcherUnrecoverableFailure"
	PrinterProcessingStateDetail_StitcherUnrecoverableStorageError         PrinterProcessingStateDetail = "stitcherUnrecoverableStorageError"
	PrinterProcessingStateDetail_StitcherWarmingUp                         PrinterProcessingStateDetail = "stitcherWarmingUp"
	PrinterProcessingStateDetail_StoppedPartially                          PrinterProcessingStateDetail = "stoppedPartially"
	PrinterProcessingStateDetail_Stopping                                  PrinterProcessingStateDetail = "stopping"
	PrinterProcessingStateDetail_SubunitAdded                              PrinterProcessingStateDetail = "subunitAdded"
	PrinterProcessingStateDetail_SubunitAlmostEmpty                        PrinterProcessingStateDetail = "subunitAlmostEmpty"
	PrinterProcessingStateDetail_SubunitAlmostFull                         PrinterProcessingStateDetail = "subunitAlmostFull"
	PrinterProcessingStateDetail_SubunitAtLimit                            PrinterProcessingStateDetail = "subunitAtLimit"
	PrinterProcessingStateDetail_SubunitClosed                             PrinterProcessingStateDetail = "subunitClosed"
	PrinterProcessingStateDetail_SubunitCoolingDown                        PrinterProcessingStateDetail = "subunitCoolingDown"
	PrinterProcessingStateDetail_SubunitEmpty                              PrinterProcessingStateDetail = "subunitEmpty"
	PrinterProcessingStateDetail_SubunitFull                               PrinterProcessingStateDetail = "subunitFull"
	PrinterProcessingStateDetail_SubunitLifeAlmostOver                     PrinterProcessingStateDetail = "subunitLifeAlmostOver"
	PrinterProcessingStateDetail_SubunitLifeOver                           PrinterProcessingStateDetail = "subunitLifeOver"
	PrinterProcessingStateDetail_SubunitMemoryExhausted                    PrinterProcessingStateDetail = "subunitMemoryExhausted"
	PrinterProcessingStateDetail_SubunitMissing                            PrinterProcessingStateDetail = "subunitMissing"
	PrinterProcessingStateDetail_SubunitMotorFailure                       PrinterProcessingStateDetail = "subunitMotorFailure"
	PrinterProcessingStateDetail_SubunitNearLimit                          PrinterProcessingStateDetail = "subunitNearLimit"
	PrinterProcessingStateDetail_SubunitOffline                            PrinterProcessingStateDetail = "subunitOffline"
	PrinterProcessingStateDetail_SubunitOpened                             PrinterProcessingStateDetail = "subunitOpened"
	PrinterProcessingStateDetail_SubunitOverTemperature                    PrinterProcessingStateDetail = "subunitOverTemperature"
	PrinterProcessingStateDetail_SubunitPowerSaver                         PrinterProcessingStateDetail = "subunitPowerSaver"
	PrinterProcessingStateDetail_SubunitRecoverableFailure                 PrinterProcessingStateDetail = "subunitRecoverableFailure"
	PrinterProcessingStateDetail_SubunitRecoverableStorage                 PrinterProcessingStateDetail = "subunitRecoverableStorage"
	PrinterProcessingStateDetail_SubunitRemoved                            PrinterProcessingStateDetail = "subunitRemoved"
	PrinterProcessingStateDetail_SubunitResourceAdded                      PrinterProcessingStateDetail = "subunitResourceAdded"
	PrinterProcessingStateDetail_SubunitResourceRemoved                    PrinterProcessingStateDetail = "subunitResourceRemoved"
	PrinterProcessingStateDetail_SubunitThermistorFailure                  PrinterProcessingStateDetail = "subunitThermistorFailure"
	PrinterProcessingStateDetail_SubunitTimingFailure                      PrinterProcessingStateDetail = "subunitTimingFailure"
	PrinterProcessingStateDetail_SubunitTurnedOff                          PrinterProcessingStateDetail = "subunitTurnedOff"
	PrinterProcessingStateDetail_SubunitTurnedOn                           PrinterProcessingStateDetail = "subunitTurnedOn"
	PrinterProcessingStateDetail_SubunitUnderTemperature                   PrinterProcessingStateDetail = "subunitUnderTemperature"
	PrinterProcessingStateDetail_SubunitUnrecoverableFailure               PrinterProcessingStateDetail = "subunitUnrecoverableFailure"
	PrinterProcessingStateDetail_SubunitUnrecoverableStorage               PrinterProcessingStateDetail = "subunitUnrecoverableStorage"
	PrinterProcessingStateDetail_SubunitWarmingUp                          PrinterProcessingStateDetail = "subunitWarmingUp"
	PrinterProcessingStateDetail_Suspend                                   PrinterProcessingStateDetail = "suspend"
	PrinterProcessingStateDetail_Testing                                   PrinterProcessingStateDetail = "testing"
	PrinterProcessingStateDetail_TimedOut                                  PrinterProcessingStateDetail = "timedOut"
	PrinterProcessingStateDetail_TonerEmpty                                PrinterProcessingStateDetail = "tonerEmpty"
	PrinterProcessingStateDetail_TonerLow                                  PrinterProcessingStateDetail = "tonerLow"
	PrinterProcessingStateDetail_TrimmerAdded                              PrinterProcessingStateDetail = "trimmerAdded"
	PrinterProcessingStateDetail_TrimmerAlmostEmpty                        PrinterProcessingStateDetail = "trimmerAlmostEmpty"
	PrinterProcessingStateDetail_TrimmerAlmostFull                         PrinterProcessingStateDetail = "trimmerAlmostFull"
	PrinterProcessingStateDetail_TrimmerAtLimit                            PrinterProcessingStateDetail = "trimmerAtLimit"
	PrinterProcessingStateDetail_TrimmerClosed                             PrinterProcessingStateDetail = "trimmerClosed"
	PrinterProcessingStateDetail_TrimmerConfigurationChange                PrinterProcessingStateDetail = "trimmerConfigurationChange"
	PrinterProcessingStateDetail_TrimmerCoverClosed                        PrinterProcessingStateDetail = "trimmerCoverClosed"
	PrinterProcessingStateDetail_TrimmerCoverOpen                          PrinterProcessingStateDetail = "trimmerCoverOpen"
	PrinterProcessingStateDetail_TrimmerEmpty                              PrinterProcessingStateDetail = "trimmerEmpty"
	PrinterProcessingStateDetail_TrimmerFull                               PrinterProcessingStateDetail = "trimmerFull"
	PrinterProcessingStateDetail_TrimmerInterlockClosed                    PrinterProcessingStateDetail = "trimmerInterlockClosed"
	PrinterProcessingStateDetail_TrimmerInterlockOpen                      PrinterProcessingStateDetail = "trimmerInterlockOpen"
	PrinterProcessingStateDetail_TrimmerJam                                PrinterProcessingStateDetail = "trimmerJam"
	PrinterProcessingStateDetail_TrimmerLifeAlmostOver                     PrinterProcessingStateDetail = "trimmerLifeAlmostOver"
	PrinterProcessingStateDetail_TrimmerLifeOver                           PrinterProcessingStateDetail = "trimmerLifeOver"
	PrinterProcessingStateDetail_TrimmerMemoryExhausted                    PrinterProcessingStateDetail = "trimmerMemoryExhausted"
	PrinterProcessingStateDetail_TrimmerMissing                            PrinterProcessingStateDetail = "trimmerMissing"
	PrinterProcessingStateDetail_TrimmerMotorFailure                       PrinterProcessingStateDetail = "trimmerMotorFailure"
	PrinterProcessingStateDetail_TrimmerNearLimit                          PrinterProcessingStateDetail = "trimmerNearLimit"
	PrinterProcessingStateDetail_TrimmerOffline                            PrinterProcessingStateDetail = "trimmerOffline"
	PrinterProcessingStateDetail_TrimmerOpened                             PrinterProcessingStateDetail = "trimmerOpened"
	PrinterProcessingStateDetail_TrimmerOverTemperature                    PrinterProcessingStateDetail = "trimmerOverTemperature"
	PrinterProcessingStateDetail_TrimmerPowerSaver                         PrinterProcessingStateDetail = "trimmerPowerSaver"
	PrinterProcessingStateDetail_TrimmerRecoverableFailure                 PrinterProcessingStateDetail = "trimmerRecoverableFailure"
	PrinterProcessingStateDetail_TrimmerRecoverableStorage                 PrinterProcessingStateDetail = "trimmerRecoverableStorage"
	PrinterProcessingStateDetail_TrimmerRemoved                            PrinterProcessingStateDetail = "trimmerRemoved"
	PrinterProcessingStateDetail_TrimmerResourceAdded                      PrinterProcessingStateDetail = "trimmerResourceAdded"
	PrinterProcessingStateDetail_TrimmerResourceRemoved                    PrinterProcessingStateDetail = "trimmerResourceRemoved"
	PrinterProcessingStateDetail_TrimmerThermistorFailure                  PrinterProcessingStateDetail = "trimmerThermistorFailure"
	PrinterProcessingStateDetail_TrimmerTimingFailure                      PrinterProcessingStateDetail = "trimmerTimingFailure"
	PrinterProcessingStateDetail_TrimmerTurnedOff                          PrinterProcessingStateDetail = "trimmerTurnedOff"
	PrinterProcessingStateDetail_TrimmerTurnedOn                           PrinterProcessingStateDetail = "trimmerTurnedOn"
	PrinterProcessingStateDetail_TrimmerUnderTemperature                   PrinterProcessingStateDetail = "trimmerUnderTemperature"
	PrinterProcessingStateDetail_TrimmerUnrecoverableFailure               PrinterProcessingStateDetail = "trimmerUnrecoverableFailure"
	PrinterProcessingStateDetail_TrimmerUnrecoverableStorageError          PrinterProcessingStateDetail = "trimmerUnrecoverableStorageError"
	PrinterProcessingStateDetail_TrimmerWarmingUp                          PrinterProcessingStateDetail = "trimmerWarmingUp"
	PrinterProcessingStateDetail_Unknown                                   PrinterProcessingStateDetail = "unknown"
	PrinterProcessingStateDetail_WrapperAdded                              PrinterProcessingStateDetail = "wrapperAdded"
	PrinterProcessingStateDetail_WrapperAlmostEmpty                        PrinterProcessingStateDetail = "wrapperAlmostEmpty"
	PrinterProcessingStateDetail_WrapperAlmostFull                         PrinterProcessingStateDetail = "wrapperAlmostFull"
	PrinterProcessingStateDetail_WrapperAtLimit                            PrinterProcessingStateDetail = "wrapperAtLimit"
	PrinterProcessingStateDetail_WrapperClosed                             PrinterProcessingStateDetail = "wrapperClosed"
	PrinterProcessingStateDetail_WrapperConfigurationChange                PrinterProcessingStateDetail = "wrapperConfigurationChange"
	PrinterProcessingStateDetail_WrapperCoverClosed                        PrinterProcessingStateDetail = "wrapperCoverClosed"
	PrinterProcessingStateDetail_WrapperCoverOpen                          PrinterProcessingStateDetail = "wrapperCoverOpen"
	PrinterProcessingStateDetail_WrapperEmpty                              PrinterProcessingStateDetail = "wrapperEmpty"
	PrinterProcessingStateDetail_WrapperFull                               PrinterProcessingStateDetail = "wrapperFull"
	PrinterProcessingStateDetail_WrapperInterlockClosed                    PrinterProcessingStateDetail = "wrapperInterlockClosed"
	PrinterProcessingStateDetail_WrapperInterlockOpen                      PrinterProcessingStateDetail = "wrapperInterlockOpen"
	PrinterProcessingStateDetail_WrapperJam                                PrinterProcessingStateDetail = "wrapperJam"
	PrinterProcessingStateDetail_WrapperLifeAlmostOver                     PrinterProcessingStateDetail = "wrapperLifeAlmostOver"
	PrinterProcessingStateDetail_WrapperLifeOver                           PrinterProcessingStateDetail = "wrapperLifeOver"
	PrinterProcessingStateDetail_WrapperMemoryExhausted                    PrinterProcessingStateDetail = "wrapperMemoryExhausted"
	PrinterProcessingStateDetail_WrapperMissing                            PrinterProcessingStateDetail = "wrapperMissing"
	PrinterProcessingStateDetail_WrapperMotorFailure                       PrinterProcessingStateDetail = "wrapperMotorFailure"
	PrinterProcessingStateDetail_WrapperNearLimit                          PrinterProcessingStateDetail = "wrapperNearLimit"
	PrinterProcessingStateDetail_WrapperOffline                            PrinterProcessingStateDetail = "wrapperOffline"
	PrinterProcessingStateDetail_WrapperOpened                             PrinterProcessingStateDetail = "wrapperOpened"
	PrinterProcessingStateDetail_WrapperOverTemperature                    PrinterProcessingStateDetail = "wrapperOverTemperature"
	PrinterProcessingStateDetail_WrapperPowerSaver                         PrinterProcessingStateDetail = "wrapperPowerSaver"
	PrinterProcessingStateDetail_WrapperRecoverableFailure                 PrinterProcessingStateDetail = "wrapperRecoverableFailure"
	PrinterProcessingStateDetail_WrapperRecoverableStorage                 PrinterProcessingStateDetail = "wrapperRecoverableStorage"
	PrinterProcessingStateDetail_WrapperRemoved                            PrinterProcessingStateDetail = "wrapperRemoved"
	PrinterProcessingStateDetail_WrapperResourceAdded                      PrinterProcessingStateDetail = "wrapperResourceAdded"
	PrinterProcessingStateDetail_WrapperResourceRemoved                    PrinterProcessingStateDetail = "wrapperResourceRemoved"
	PrinterProcessingStateDetail_WrapperThermistorFailure                  PrinterProcessingStateDetail = "wrapperThermistorFailure"
	PrinterProcessingStateDetail_WrapperTimingFailure                      PrinterProcessingStateDetail = "wrapperTimingFailure"
	PrinterProcessingStateDetail_WrapperTurnedOff                          PrinterProcessingStateDetail = "wrapperTurnedOff"
	PrinterProcessingStateDetail_WrapperTurnedOn                           PrinterProcessingStateDetail = "wrapperTurnedOn"
	PrinterProcessingStateDetail_WrapperUnderTemperature                   PrinterProcessingStateDetail = "wrapperUnderTemperature"
	PrinterProcessingStateDetail_WrapperUnrecoverableFailure               PrinterProcessingStateDetail = "wrapperUnrecoverableFailure"
	PrinterProcessingStateDetail_WrapperUnrecoverableStorageError          PrinterProcessingStateDetail = "wrapperUnrecoverableStorageError"
	PrinterProcessingStateDetail_WrapperWarmingUp                          PrinterProcessingStateDetail = "wrapperWarmingUp"
)

func PossibleValuesForPrinterProcessingStateDetail() []string {
	return []string{
		string(PrinterProcessingStateDetail_AlertRemovalOfBinaryChangeEntry),
		string(PrinterProcessingStateDetail_BanderAdded),
		string(PrinterProcessingStateDetail_BanderAlmostEmpty),
		string(PrinterProcessingStateDetail_BanderAlmostFull),
		string(PrinterProcessingStateDetail_BanderAtLimit),
		string(PrinterProcessingStateDetail_BanderClosed),
		string(PrinterProcessingStateDetail_BanderConfigurationChange),
		string(PrinterProcessingStateDetail_BanderCoverClosed),
		string(PrinterProcessingStateDetail_BanderCoverOpen),
		string(PrinterProcessingStateDetail_BanderEmpty),
		string(PrinterProcessingStateDetail_BanderFull),
		string(PrinterProcessingStateDetail_BanderInterlockClosed),
		string(PrinterProcessingStateDetail_BanderInterlockOpen),
		string(PrinterProcessingStateDetail_BanderJam),
		string(PrinterProcessingStateDetail_BanderLifeAlmostOver),
		string(PrinterProcessingStateDetail_BanderLifeOver),
		string(PrinterProcessingStateDetail_BanderMemoryExhausted),
		string(PrinterProcessingStateDetail_BanderMissing),
		string(PrinterProcessingStateDetail_BanderMotorFailure),
		string(PrinterProcessingStateDetail_BanderNearLimit),
		string(PrinterProcessingStateDetail_BanderOffline),
		string(PrinterProcessingStateDetail_BanderOpened),
		string(PrinterProcessingStateDetail_BanderOverTemperature),
		string(PrinterProcessingStateDetail_BanderPowerSaver),
		string(PrinterProcessingStateDetail_BanderRecoverableFailure),
		string(PrinterProcessingStateDetail_BanderRecoverableStorage),
		string(PrinterProcessingStateDetail_BanderRemoved),
		string(PrinterProcessingStateDetail_BanderResourceAdded),
		string(PrinterProcessingStateDetail_BanderResourceRemoved),
		string(PrinterProcessingStateDetail_BanderThermistorFailure),
		string(PrinterProcessingStateDetail_BanderTimingFailure),
		string(PrinterProcessingStateDetail_BanderTurnedOff),
		string(PrinterProcessingStateDetail_BanderTurnedOn),
		string(PrinterProcessingStateDetail_BanderUnderTemperature),
		string(PrinterProcessingStateDetail_BanderUnrecoverableFailure),
		string(PrinterProcessingStateDetail_BanderUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_BanderWarmingUp),
		string(PrinterProcessingStateDetail_BinderAdded),
		string(PrinterProcessingStateDetail_BinderAlmostEmpty),
		string(PrinterProcessingStateDetail_BinderAlmostFull),
		string(PrinterProcessingStateDetail_BinderAtLimit),
		string(PrinterProcessingStateDetail_BinderClosed),
		string(PrinterProcessingStateDetail_BinderConfigurationChange),
		string(PrinterProcessingStateDetail_BinderCoverClosed),
		string(PrinterProcessingStateDetail_BinderCoverOpen),
		string(PrinterProcessingStateDetail_BinderEmpty),
		string(PrinterProcessingStateDetail_BinderFull),
		string(PrinterProcessingStateDetail_BinderInterlockClosed),
		string(PrinterProcessingStateDetail_BinderInterlockOpen),
		string(PrinterProcessingStateDetail_BinderJam),
		string(PrinterProcessingStateDetail_BinderLifeAlmostOver),
		string(PrinterProcessingStateDetail_BinderLifeOver),
		string(PrinterProcessingStateDetail_BinderMemoryExhausted),
		string(PrinterProcessingStateDetail_BinderMissing),
		string(PrinterProcessingStateDetail_BinderMotorFailure),
		string(PrinterProcessingStateDetail_BinderNearLimit),
		string(PrinterProcessingStateDetail_BinderOffline),
		string(PrinterProcessingStateDetail_BinderOpened),
		string(PrinterProcessingStateDetail_BinderOverTemperature),
		string(PrinterProcessingStateDetail_BinderPowerSaver),
		string(PrinterProcessingStateDetail_BinderRecoverableFailure),
		string(PrinterProcessingStateDetail_BinderRecoverableStorage),
		string(PrinterProcessingStateDetail_BinderRemoved),
		string(PrinterProcessingStateDetail_BinderResourceAdded),
		string(PrinterProcessingStateDetail_BinderResourceRemoved),
		string(PrinterProcessingStateDetail_BinderThermistorFailure),
		string(PrinterProcessingStateDetail_BinderTimingFailure),
		string(PrinterProcessingStateDetail_BinderTurnedOff),
		string(PrinterProcessingStateDetail_BinderTurnedOn),
		string(PrinterProcessingStateDetail_BinderUnderTemperature),
		string(PrinterProcessingStateDetail_BinderUnrecoverableFailure),
		string(PrinterProcessingStateDetail_BinderUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_BinderWarmingUp),
		string(PrinterProcessingStateDetail_CameraFailure),
		string(PrinterProcessingStateDetail_ChamberCooling),
		string(PrinterProcessingStateDetail_ChamberFailure),
		string(PrinterProcessingStateDetail_ChamberHeating),
		string(PrinterProcessingStateDetail_ChamberTemperatureHigh),
		string(PrinterProcessingStateDetail_ChamberTemperatureLow),
		string(PrinterProcessingStateDetail_CleanerLifeAlmostOver),
		string(PrinterProcessingStateDetail_CleanerLifeOver),
		string(PrinterProcessingStateDetail_ConfigurationChange),
		string(PrinterProcessingStateDetail_ConnectingToDevice),
		string(PrinterProcessingStateDetail_CoverOpen),
		string(PrinterProcessingStateDetail_Deactivated),
		string(PrinterProcessingStateDetail_Deleted),
		string(PrinterProcessingStateDetail_DeveloperEmpty),
		string(PrinterProcessingStateDetail_DeveloperLow),
		string(PrinterProcessingStateDetail_DieCutterAdded),
		string(PrinterProcessingStateDetail_DieCutterAlmostEmpty),
		string(PrinterProcessingStateDetail_DieCutterAlmostFull),
		string(PrinterProcessingStateDetail_DieCutterAtLimit),
		string(PrinterProcessingStateDetail_DieCutterClosed),
		string(PrinterProcessingStateDetail_DieCutterConfigurationChange),
		string(PrinterProcessingStateDetail_DieCutterCoverClosed),
		string(PrinterProcessingStateDetail_DieCutterCoverOpen),
		string(PrinterProcessingStateDetail_DieCutterEmpty),
		string(PrinterProcessingStateDetail_DieCutterFull),
		string(PrinterProcessingStateDetail_DieCutterInterlockClosed),
		string(PrinterProcessingStateDetail_DieCutterInterlockOpen),
		string(PrinterProcessingStateDetail_DieCutterJam),
		string(PrinterProcessingStateDetail_DieCutterLifeAlmostOver),
		string(PrinterProcessingStateDetail_DieCutterLifeOver),
		string(PrinterProcessingStateDetail_DieCutterMemoryExhausted),
		string(PrinterProcessingStateDetail_DieCutterMissing),
		string(PrinterProcessingStateDetail_DieCutterMotorFailure),
		string(PrinterProcessingStateDetail_DieCutterNearLimit),
		string(PrinterProcessingStateDetail_DieCutterOffline),
		string(PrinterProcessingStateDetail_DieCutterOpened),
		string(PrinterProcessingStateDetail_DieCutterOverTemperature),
		string(PrinterProcessingStateDetail_DieCutterPowerSaver),
		string(PrinterProcessingStateDetail_DieCutterRecoverableFailure),
		string(PrinterProcessingStateDetail_DieCutterRecoverableStorage),
		string(PrinterProcessingStateDetail_DieCutterRemoved),
		string(PrinterProcessingStateDetail_DieCutterResourceAdded),
		string(PrinterProcessingStateDetail_DieCutterResourceRemoved),
		string(PrinterProcessingStateDetail_DieCutterThermistorFailure),
		string(PrinterProcessingStateDetail_DieCutterTimingFailure),
		string(PrinterProcessingStateDetail_DieCutterTurnedOff),
		string(PrinterProcessingStateDetail_DieCutterTurnedOn),
		string(PrinterProcessingStateDetail_DieCutterUnderTemperature),
		string(PrinterProcessingStateDetail_DieCutterUnrecoverableFailure),
		string(PrinterProcessingStateDetail_DieCutterUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_DieCutterWarmingUp),
		string(PrinterProcessingStateDetail_DoorOpen),
		string(PrinterProcessingStateDetail_ExtruderCooling),
		string(PrinterProcessingStateDetail_ExtruderFailure),
		string(PrinterProcessingStateDetail_ExtruderHeating),
		string(PrinterProcessingStateDetail_ExtruderJam),
		string(PrinterProcessingStateDetail_ExtruderTemperatureHigh),
		string(PrinterProcessingStateDetail_ExtruderTemperatureLow),
		string(PrinterProcessingStateDetail_FanFailure),
		string(PrinterProcessingStateDetail_FaxModemLifeAlmostOver),
		string(PrinterProcessingStateDetail_FaxModemLifeOver),
		string(PrinterProcessingStateDetail_FaxModemMissing),
		string(PrinterProcessingStateDetail_FaxModemTurnedOff),
		string(PrinterProcessingStateDetail_FaxModemTurnedOn),
		string(PrinterProcessingStateDetail_FolderAdded),
		string(PrinterProcessingStateDetail_FolderAlmostEmpty),
		string(PrinterProcessingStateDetail_FolderAlmostFull),
		string(PrinterProcessingStateDetail_FolderAtLimit),
		string(PrinterProcessingStateDetail_FolderClosed),
		string(PrinterProcessingStateDetail_FolderConfigurationChange),
		string(PrinterProcessingStateDetail_FolderCoverClosed),
		string(PrinterProcessingStateDetail_FolderCoverOpen),
		string(PrinterProcessingStateDetail_FolderEmpty),
		string(PrinterProcessingStateDetail_FolderFull),
		string(PrinterProcessingStateDetail_FolderInterlockClosed),
		string(PrinterProcessingStateDetail_FolderInterlockOpen),
		string(PrinterProcessingStateDetail_FolderJam),
		string(PrinterProcessingStateDetail_FolderLifeAlmostOver),
		string(PrinterProcessingStateDetail_FolderLifeOver),
		string(PrinterProcessingStateDetail_FolderMemoryExhausted),
		string(PrinterProcessingStateDetail_FolderMissing),
		string(PrinterProcessingStateDetail_FolderMotorFailure),
		string(PrinterProcessingStateDetail_FolderNearLimit),
		string(PrinterProcessingStateDetail_FolderOffline),
		string(PrinterProcessingStateDetail_FolderOpened),
		string(PrinterProcessingStateDetail_FolderOverTemperature),
		string(PrinterProcessingStateDetail_FolderPowerSaver),
		string(PrinterProcessingStateDetail_FolderRecoverableFailure),
		string(PrinterProcessingStateDetail_FolderRecoverableStorage),
		string(PrinterProcessingStateDetail_FolderRemoved),
		string(PrinterProcessingStateDetail_FolderResourceAdded),
		string(PrinterProcessingStateDetail_FolderResourceRemoved),
		string(PrinterProcessingStateDetail_FolderThermistorFailure),
		string(PrinterProcessingStateDetail_FolderTimingFailure),
		string(PrinterProcessingStateDetail_FolderTurnedOff),
		string(PrinterProcessingStateDetail_FolderTurnedOn),
		string(PrinterProcessingStateDetail_FolderUnderTemperature),
		string(PrinterProcessingStateDetail_FolderUnrecoverableFailure),
		string(PrinterProcessingStateDetail_FolderUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_FolderWarmingUp),
		string(PrinterProcessingStateDetail_FuserOverTemp),
		string(PrinterProcessingStateDetail_FuserUnderTemp),
		string(PrinterProcessingStateDetail_Hibernate),
		string(PrinterProcessingStateDetail_HoldNewJobs),
		string(PrinterProcessingStateDetail_IdentifyPrinterRequested),
		string(PrinterProcessingStateDetail_ImprinterAdded),
		string(PrinterProcessingStateDetail_ImprinterAlmostEmpty),
		string(PrinterProcessingStateDetail_ImprinterAlmostFull),
		string(PrinterProcessingStateDetail_ImprinterAtLimit),
		string(PrinterProcessingStateDetail_ImprinterClosed),
		string(PrinterProcessingStateDetail_ImprinterConfigurationChange),
		string(PrinterProcessingStateDetail_ImprinterCoverClosed),
		string(PrinterProcessingStateDetail_ImprinterCoverOpen),
		string(PrinterProcessingStateDetail_ImprinterEmpty),
		string(PrinterProcessingStateDetail_ImprinterFull),
		string(PrinterProcessingStateDetail_ImprinterInterlockClosed),
		string(PrinterProcessingStateDetail_ImprinterInterlockOpen),
		string(PrinterProcessingStateDetail_ImprinterJam),
		string(PrinterProcessingStateDetail_ImprinterLifeAlmostOver),
		string(PrinterProcessingStateDetail_ImprinterLifeOver),
		string(PrinterProcessingStateDetail_ImprinterMemoryExhausted),
		string(PrinterProcessingStateDetail_ImprinterMissing),
		string(PrinterProcessingStateDetail_ImprinterMotorFailure),
		string(PrinterProcessingStateDetail_ImprinterNearLimit),
		string(PrinterProcessingStateDetail_ImprinterOffline),
		string(PrinterProcessingStateDetail_ImprinterOpened),
		string(PrinterProcessingStateDetail_ImprinterOverTemperature),
		string(PrinterProcessingStateDetail_ImprinterPowerSaver),
		string(PrinterProcessingStateDetail_ImprinterRecoverableFailure),
		string(PrinterProcessingStateDetail_ImprinterRecoverableStorage),
		string(PrinterProcessingStateDetail_ImprinterRemoved),
		string(PrinterProcessingStateDetail_ImprinterResourceAdded),
		string(PrinterProcessingStateDetail_ImprinterResourceRemoved),
		string(PrinterProcessingStateDetail_ImprinterThermistorFailure),
		string(PrinterProcessingStateDetail_ImprinterTimingFailure),
		string(PrinterProcessingStateDetail_ImprinterTurnedOff),
		string(PrinterProcessingStateDetail_ImprinterTurnedOn),
		string(PrinterProcessingStateDetail_ImprinterUnderTemperature),
		string(PrinterProcessingStateDetail_ImprinterUnrecoverableFailure),
		string(PrinterProcessingStateDetail_ImprinterUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_ImprinterWarmingUp),
		string(PrinterProcessingStateDetail_InputCannotFeedSizeSelected),
		string(PrinterProcessingStateDetail_InputManualInputRequest),
		string(PrinterProcessingStateDetail_InputMediaColorChange),
		string(PrinterProcessingStateDetail_InputMediaFormPartsChange),
		string(PrinterProcessingStateDetail_InputMediaSizeChange),
		string(PrinterProcessingStateDetail_InputMediaTrayFailure),
		string(PrinterProcessingStateDetail_InputMediaTrayFeedError),
		string(PrinterProcessingStateDetail_InputMediaTrayJam),
		string(PrinterProcessingStateDetail_InputMediaTypeChange),
		string(PrinterProcessingStateDetail_InputMediaWeightChange),
		string(PrinterProcessingStateDetail_InputPickRollerFailure),
		string(PrinterProcessingStateDetail_InputPickRollerLifeOver),
		string(PrinterProcessingStateDetail_InputPickRollerLifeWarn),
		string(PrinterProcessingStateDetail_InputPickRollerMissing),
		string(PrinterProcessingStateDetail_InputTrayElevationFailure),
		string(PrinterProcessingStateDetail_InputTrayMissing),
		string(PrinterProcessingStateDetail_InputTrayPositionFailure),
		string(PrinterProcessingStateDetail_InserterAdded),
		string(PrinterProcessingStateDetail_InserterAlmostEmpty),
		string(PrinterProcessingStateDetail_InserterAlmostFull),
		string(PrinterProcessingStateDetail_InserterAtLimit),
		string(PrinterProcessingStateDetail_InserterClosed),
		string(PrinterProcessingStateDetail_InserterConfigurationChange),
		string(PrinterProcessingStateDetail_InserterCoverClosed),
		string(PrinterProcessingStateDetail_InserterCoverOpen),
		string(PrinterProcessingStateDetail_InserterEmpty),
		string(PrinterProcessingStateDetail_InserterFull),
		string(PrinterProcessingStateDetail_InserterInterlockClosed),
		string(PrinterProcessingStateDetail_InserterInterlockOpen),
		string(PrinterProcessingStateDetail_InserterJam),
		string(PrinterProcessingStateDetail_InserterLifeAlmostOver),
		string(PrinterProcessingStateDetail_InserterLifeOver),
		string(PrinterProcessingStateDetail_InserterMemoryExhausted),
		string(PrinterProcessingStateDetail_InserterMissing),
		string(PrinterProcessingStateDetail_InserterMotorFailure),
		string(PrinterProcessingStateDetail_InserterNearLimit),
		string(PrinterProcessingStateDetail_InserterOffline),
		string(PrinterProcessingStateDetail_InserterOpened),
		string(PrinterProcessingStateDetail_InserterOverTemperature),
		string(PrinterProcessingStateDetail_InserterPowerSaver),
		string(PrinterProcessingStateDetail_InserterRecoverableFailure),
		string(PrinterProcessingStateDetail_InserterRecoverableStorage),
		string(PrinterProcessingStateDetail_InserterRemoved),
		string(PrinterProcessingStateDetail_InserterResourceAdded),
		string(PrinterProcessingStateDetail_InserterResourceRemoved),
		string(PrinterProcessingStateDetail_InserterThermistorFailure),
		string(PrinterProcessingStateDetail_InserterTimingFailure),
		string(PrinterProcessingStateDetail_InserterTurnedOff),
		string(PrinterProcessingStateDetail_InserterTurnedOn),
		string(PrinterProcessingStateDetail_InserterUnderTemperature),
		string(PrinterProcessingStateDetail_InserterUnrecoverableFailure),
		string(PrinterProcessingStateDetail_InserterUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_InserterWarmingUp),
		string(PrinterProcessingStateDetail_InterlockClosed),
		string(PrinterProcessingStateDetail_InterlockOpen),
		string(PrinterProcessingStateDetail_InterpreterCartridgeAdded),
		string(PrinterProcessingStateDetail_InterpreterCartridgeDeleted),
		string(PrinterProcessingStateDetail_InterpreterComplexPageEncountered),
		string(PrinterProcessingStateDetail_InterpreterMemoryDecrease),
		string(PrinterProcessingStateDetail_InterpreterMemoryIncrease),
		string(PrinterProcessingStateDetail_InterpreterResourceAdded),
		string(PrinterProcessingStateDetail_InterpreterResourceDeleted),
		string(PrinterProcessingStateDetail_InterpreterResourceUnavailable),
		string(PrinterProcessingStateDetail_LampAtEol),
		string(PrinterProcessingStateDetail_LampFailure),
		string(PrinterProcessingStateDetail_LampNearEol),
		string(PrinterProcessingStateDetail_LaserAtEol),
		string(PrinterProcessingStateDetail_LaserFailure),
		string(PrinterProcessingStateDetail_LaserNearEol),
		string(PrinterProcessingStateDetail_MakeEnvelopeAdded),
		string(PrinterProcessingStateDetail_MakeEnvelopeAlmostEmpty),
		string(PrinterProcessingStateDetail_MakeEnvelopeAlmostFull),
		string(PrinterProcessingStateDetail_MakeEnvelopeAtLimit),
		string(PrinterProcessingStateDetail_MakeEnvelopeClosed),
		string(PrinterProcessingStateDetail_MakeEnvelopeConfigurationChange),
		string(PrinterProcessingStateDetail_MakeEnvelopeCoverClosed),
		string(PrinterProcessingStateDetail_MakeEnvelopeCoverOpen),
		string(PrinterProcessingStateDetail_MakeEnvelopeEmpty),
		string(PrinterProcessingStateDetail_MakeEnvelopeFull),
		string(PrinterProcessingStateDetail_MakeEnvelopeInterlockClosed),
		string(PrinterProcessingStateDetail_MakeEnvelopeInterlockOpen),
		string(PrinterProcessingStateDetail_MakeEnvelopeJam),
		string(PrinterProcessingStateDetail_MakeEnvelopeLifeAlmostOver),
		string(PrinterProcessingStateDetail_MakeEnvelopeLifeOver),
		string(PrinterProcessingStateDetail_MakeEnvelopeMemoryExhausted),
		string(PrinterProcessingStateDetail_MakeEnvelopeMissing),
		string(PrinterProcessingStateDetail_MakeEnvelopeMotorFailure),
		string(PrinterProcessingStateDetail_MakeEnvelopeNearLimit),
		string(PrinterProcessingStateDetail_MakeEnvelopeOffline),
		string(PrinterProcessingStateDetail_MakeEnvelopeOpened),
		string(PrinterProcessingStateDetail_MakeEnvelopeOverTemperature),
		string(PrinterProcessingStateDetail_MakeEnvelopePowerSaver),
		string(PrinterProcessingStateDetail_MakeEnvelopeRecoverableFailure),
		string(PrinterProcessingStateDetail_MakeEnvelopeRecoverableStorage),
		string(PrinterProcessingStateDetail_MakeEnvelopeRemoved),
		string(PrinterProcessingStateDetail_MakeEnvelopeResourceAdded),
		string(PrinterProcessingStateDetail_MakeEnvelopeResourceRemoved),
		string(PrinterProcessingStateDetail_MakeEnvelopeThermistorFailure),
		string(PrinterProcessingStateDetail_MakeEnvelopeTimingFailure),
		string(PrinterProcessingStateDetail_MakeEnvelopeTurnedOff),
		string(PrinterProcessingStateDetail_MakeEnvelopeTurnedOn),
		string(PrinterProcessingStateDetail_MakeEnvelopeUnderTemperature),
		string(PrinterProcessingStateDetail_MakeEnvelopeUnrecoverableFailure),
		string(PrinterProcessingStateDetail_MakeEnvelopeUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_MakeEnvelopeWarmingUp),
		string(PrinterProcessingStateDetail_MarkerAdjustingPrintQuality),
		string(PrinterProcessingStateDetail_MarkerCleanerMissing),
		string(PrinterProcessingStateDetail_MarkerDeveloperAlmostEmpty),
		string(PrinterProcessingStateDetail_MarkerDeveloperEmpty),
		string(PrinterProcessingStateDetail_MarkerDeveloperMissing),
		string(PrinterProcessingStateDetail_MarkerFuserMissing),
		string(PrinterProcessingStateDetail_MarkerFuserThermistorFailure),
		string(PrinterProcessingStateDetail_MarkerFuserTimingFailure),
		string(PrinterProcessingStateDetail_MarkerInkAlmostEmpty),
		string(PrinterProcessingStateDetail_MarkerInkEmpty),
		string(PrinterProcessingStateDetail_MarkerInkMissing),
		string(PrinterProcessingStateDetail_MarkerOpcMissing),
		string(PrinterProcessingStateDetail_MarkerPrintRibbonAlmostEmpty),
		string(PrinterProcessingStateDetail_MarkerPrintRibbonEmpty),
		string(PrinterProcessingStateDetail_MarkerPrintRibbonMissing),
		string(PrinterProcessingStateDetail_MarkerSupplyAlmostEmpty),
		string(PrinterProcessingStateDetail_MarkerSupplyEmpty),
		string(PrinterProcessingStateDetail_MarkerSupplyLow),
		string(PrinterProcessingStateDetail_MarkerSupplyMissing),
		string(PrinterProcessingStateDetail_MarkerTonerCartridgeMissing),
		string(PrinterProcessingStateDetail_MarkerTonerMissing),
		string(PrinterProcessingStateDetail_MarkerWasteAlmostFull),
		string(PrinterProcessingStateDetail_MarkerWasteFull),
		string(PrinterProcessingStateDetail_MarkerWasteInkReceptacleAlmostFull),
		string(PrinterProcessingStateDetail_MarkerWasteInkReceptacleFull),
		string(PrinterProcessingStateDetail_MarkerWasteInkReceptacleMissing),
		string(PrinterProcessingStateDetail_MarkerWasteMissing),
		string(PrinterProcessingStateDetail_MarkerWasteTonerReceptacleAlmostFull),
		string(PrinterProcessingStateDetail_MarkerWasteTonerReceptacleFull),
		string(PrinterProcessingStateDetail_MarkerWasteTonerReceptacleMissing),
		string(PrinterProcessingStateDetail_MaterialEmpty),
		string(PrinterProcessingStateDetail_MaterialLow),
		string(PrinterProcessingStateDetail_MaterialNeeded),
		string(PrinterProcessingStateDetail_MediaDrying),
		string(PrinterProcessingStateDetail_MediaEmpty),
		string(PrinterProcessingStateDetail_MediaJam),
		string(PrinterProcessingStateDetail_MediaLow),
		string(PrinterProcessingStateDetail_MediaNeeded),
		string(PrinterProcessingStateDetail_MediaPathCannotDuplexMediaSelected),
		string(PrinterProcessingStateDetail_MediaPathFailure),
		string(PrinterProcessingStateDetail_MediaPathInputEmpty),
		string(PrinterProcessingStateDetail_MediaPathInputFeedError),
		string(PrinterProcessingStateDetail_MediaPathInputJam),
		string(PrinterProcessingStateDetail_MediaPathInputRequest),
		string(PrinterProcessingStateDetail_MediaPathJam),
		string(PrinterProcessingStateDetail_MediaPathMediaTrayAlmostFull),
		string(PrinterProcessingStateDetail_MediaPathMediaTrayFull),
		string(PrinterProcessingStateDetail_MediaPathMediaTrayMissing),
		string(PrinterProcessingStateDetail_MediaPathOutputFeedError),
		string(PrinterProcessingStateDetail_MediaPathOutputFull),
		string(PrinterProcessingStateDetail_MediaPathOutputJam),
		string(PrinterProcessingStateDetail_MediaPathPickRollerFailure),
		string(PrinterProcessingStateDetail_MediaPathPickRollerLifeOver),
		string(PrinterProcessingStateDetail_MediaPathPickRollerLifeWarn),
		string(PrinterProcessingStateDetail_MediaPathPickRollerMissing),
		string(PrinterProcessingStateDetail_MotorFailure),
		string(PrinterProcessingStateDetail_MovingToPaused),
		string(PrinterProcessingStateDetail_None),
		string(PrinterProcessingStateDetail_OpticalPhotoConductorLifeOver),
		string(PrinterProcessingStateDetail_OpticalPhotoConductorNearEndOfLife),
		string(PrinterProcessingStateDetail_Other),
		string(PrinterProcessingStateDetail_OutputAreaAlmostFull),
		string(PrinterProcessingStateDetail_OutputAreaFull),
		string(PrinterProcessingStateDetail_OutputMailboxSelectFailure),
		string(PrinterProcessingStateDetail_OutputMediaTrayFailure),
		string(PrinterProcessingStateDetail_OutputMediaTrayFeedError),
		string(PrinterProcessingStateDetail_OutputMediaTrayJam),
		string(PrinterProcessingStateDetail_OutputTrayMissing),
		string(PrinterProcessingStateDetail_Paused),
		string(PrinterProcessingStateDetail_PerforaterAdded),
		string(PrinterProcessingStateDetail_PerforaterAlmostEmpty),
		string(PrinterProcessingStateDetail_PerforaterAlmostFull),
		string(PrinterProcessingStateDetail_PerforaterAtLimit),
		string(PrinterProcessingStateDetail_PerforaterClosed),
		string(PrinterProcessingStateDetail_PerforaterConfigurationChange),
		string(PrinterProcessingStateDetail_PerforaterCoverClosed),
		string(PrinterProcessingStateDetail_PerforaterCoverOpen),
		string(PrinterProcessingStateDetail_PerforaterEmpty),
		string(PrinterProcessingStateDetail_PerforaterFull),
		string(PrinterProcessingStateDetail_PerforaterInterlockClosed),
		string(PrinterProcessingStateDetail_PerforaterInterlockOpen),
		string(PrinterProcessingStateDetail_PerforaterJam),
		string(PrinterProcessingStateDetail_PerforaterLifeAlmostOver),
		string(PrinterProcessingStateDetail_PerforaterLifeOver),
		string(PrinterProcessingStateDetail_PerforaterMemoryExhausted),
		string(PrinterProcessingStateDetail_PerforaterMissing),
		string(PrinterProcessingStateDetail_PerforaterMotorFailure),
		string(PrinterProcessingStateDetail_PerforaterNearLimit),
		string(PrinterProcessingStateDetail_PerforaterOffline),
		string(PrinterProcessingStateDetail_PerforaterOpened),
		string(PrinterProcessingStateDetail_PerforaterOverTemperature),
		string(PrinterProcessingStateDetail_PerforaterPowerSaver),
		string(PrinterProcessingStateDetail_PerforaterRecoverableFailure),
		string(PrinterProcessingStateDetail_PerforaterRecoverableStorage),
		string(PrinterProcessingStateDetail_PerforaterRemoved),
		string(PrinterProcessingStateDetail_PerforaterResourceAdded),
		string(PrinterProcessingStateDetail_PerforaterResourceRemoved),
		string(PrinterProcessingStateDetail_PerforaterThermistorFailure),
		string(PrinterProcessingStateDetail_PerforaterTimingFailure),
		string(PrinterProcessingStateDetail_PerforaterTurnedOff),
		string(PrinterProcessingStateDetail_PerforaterTurnedOn),
		string(PrinterProcessingStateDetail_PerforaterUnderTemperature),
		string(PrinterProcessingStateDetail_PerforaterUnrecoverableFailure),
		string(PrinterProcessingStateDetail_PerforaterUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_PerforaterWarmingUp),
		string(PrinterProcessingStateDetail_PlatformCooling),
		string(PrinterProcessingStateDetail_PlatformFailure),
		string(PrinterProcessingStateDetail_PlatformHeating),
		string(PrinterProcessingStateDetail_PlatformTemperatureHigh),
		string(PrinterProcessingStateDetail_PlatformTemperatureLow),
		string(PrinterProcessingStateDetail_PowerDown),
		string(PrinterProcessingStateDetail_PowerUp),
		string(PrinterProcessingStateDetail_PrinterManualReset),
		string(PrinterProcessingStateDetail_PrinterNmsReset),
		string(PrinterProcessingStateDetail_PrinterReadyToPrint),
		string(PrinterProcessingStateDetail_PuncherAdded),
		string(PrinterProcessingStateDetail_PuncherAlmostEmpty),
		string(PrinterProcessingStateDetail_PuncherAlmostFull),
		string(PrinterProcessingStateDetail_PuncherAtLimit),
		string(PrinterProcessingStateDetail_PuncherClosed),
		string(PrinterProcessingStateDetail_PuncherConfigurationChange),
		string(PrinterProcessingStateDetail_PuncherCoverClosed),
		string(PrinterProcessingStateDetail_PuncherCoverOpen),
		string(PrinterProcessingStateDetail_PuncherEmpty),
		string(PrinterProcessingStateDetail_PuncherFull),
		string(PrinterProcessingStateDetail_PuncherInterlockClosed),
		string(PrinterProcessingStateDetail_PuncherInterlockOpen),
		string(PrinterProcessingStateDetail_PuncherJam),
		string(PrinterProcessingStateDetail_PuncherLifeAlmostOver),
		string(PrinterProcessingStateDetail_PuncherLifeOver),
		string(PrinterProcessingStateDetail_PuncherMemoryExhausted),
		string(PrinterProcessingStateDetail_PuncherMissing),
		string(PrinterProcessingStateDetail_PuncherMotorFailure),
		string(PrinterProcessingStateDetail_PuncherNearLimit),
		string(PrinterProcessingStateDetail_PuncherOffline),
		string(PrinterProcessingStateDetail_PuncherOpened),
		string(PrinterProcessingStateDetail_PuncherOverTemperature),
		string(PrinterProcessingStateDetail_PuncherPowerSaver),
		string(PrinterProcessingStateDetail_PuncherRecoverableFailure),
		string(PrinterProcessingStateDetail_PuncherRecoverableStorage),
		string(PrinterProcessingStateDetail_PuncherRemoved),
		string(PrinterProcessingStateDetail_PuncherResourceAdded),
		string(PrinterProcessingStateDetail_PuncherResourceRemoved),
		string(PrinterProcessingStateDetail_PuncherThermistorFailure),
		string(PrinterProcessingStateDetail_PuncherTimingFailure),
		string(PrinterProcessingStateDetail_PuncherTurnedOff),
		string(PrinterProcessingStateDetail_PuncherTurnedOn),
		string(PrinterProcessingStateDetail_PuncherUnderTemperature),
		string(PrinterProcessingStateDetail_PuncherUnrecoverableFailure),
		string(PrinterProcessingStateDetail_PuncherUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_PuncherWarmingUp),
		string(PrinterProcessingStateDetail_Resuming),
		string(PrinterProcessingStateDetail_ScanMediaPathFailure),
		string(PrinterProcessingStateDetail_ScanMediaPathInputEmpty),
		string(PrinterProcessingStateDetail_ScanMediaPathInputFeedError),
		string(PrinterProcessingStateDetail_ScanMediaPathInputJam),
		string(PrinterProcessingStateDetail_ScanMediaPathInputRequest),
		string(PrinterProcessingStateDetail_ScanMediaPathJam),
		string(PrinterProcessingStateDetail_ScanMediaPathOutputFeedError),
		string(PrinterProcessingStateDetail_ScanMediaPathOutputFull),
		string(PrinterProcessingStateDetail_ScanMediaPathOutputJam),
		string(PrinterProcessingStateDetail_ScanMediaPathPickRollerFailure),
		string(PrinterProcessingStateDetail_ScanMediaPathPickRollerLifeOver),
		string(PrinterProcessingStateDetail_ScanMediaPathPickRollerLifeWarn),
		string(PrinterProcessingStateDetail_ScanMediaPathPickRollerMissing),
		string(PrinterProcessingStateDetail_ScanMediaPathTrayAlmostFull),
		string(PrinterProcessingStateDetail_ScanMediaPathTrayFull),
		string(PrinterProcessingStateDetail_ScanMediaPathTrayMissing),
		string(PrinterProcessingStateDetail_ScannerLightFailure),
		string(PrinterProcessingStateDetail_ScannerLightLifeAlmostOver),
		string(PrinterProcessingStateDetail_ScannerLightLifeOver),
		string(PrinterProcessingStateDetail_ScannerLightMissing),
		string(PrinterProcessingStateDetail_ScannerSensorFailure),
		string(PrinterProcessingStateDetail_ScannerSensorLifeAlmostOver),
		string(PrinterProcessingStateDetail_ScannerSensorLifeOver),
		string(PrinterProcessingStateDetail_ScannerSensorMissing),
		string(PrinterProcessingStateDetail_SeparationCutterAdded),
		string(PrinterProcessingStateDetail_SeparationCutterAlmostEmpty),
		string(PrinterProcessingStateDetail_SeparationCutterAlmostFull),
		string(PrinterProcessingStateDetail_SeparationCutterAtLimit),
		string(PrinterProcessingStateDetail_SeparationCutterClosed),
		string(PrinterProcessingStateDetail_SeparationCutterConfigurationChange),
		string(PrinterProcessingStateDetail_SeparationCutterCoverClosed),
		string(PrinterProcessingStateDetail_SeparationCutterCoverOpen),
		string(PrinterProcessingStateDetail_SeparationCutterEmpty),
		string(PrinterProcessingStateDetail_SeparationCutterFull),
		string(PrinterProcessingStateDetail_SeparationCutterInterlockClosed),
		string(PrinterProcessingStateDetail_SeparationCutterInterlockOpen),
		string(PrinterProcessingStateDetail_SeparationCutterJam),
		string(PrinterProcessingStateDetail_SeparationCutterLifeAlmostOver),
		string(PrinterProcessingStateDetail_SeparationCutterLifeOver),
		string(PrinterProcessingStateDetail_SeparationCutterMemoryExhausted),
		string(PrinterProcessingStateDetail_SeparationCutterMissing),
		string(PrinterProcessingStateDetail_SeparationCutterMotorFailure),
		string(PrinterProcessingStateDetail_SeparationCutterNearLimit),
		string(PrinterProcessingStateDetail_SeparationCutterOffline),
		string(PrinterProcessingStateDetail_SeparationCutterOpened),
		string(PrinterProcessingStateDetail_SeparationCutterOverTemperature),
		string(PrinterProcessingStateDetail_SeparationCutterPowerSaver),
		string(PrinterProcessingStateDetail_SeparationCutterRecoverableFailure),
		string(PrinterProcessingStateDetail_SeparationCutterRecoverableStorage),
		string(PrinterProcessingStateDetail_SeparationCutterRemoved),
		string(PrinterProcessingStateDetail_SeparationCutterResourceAdded),
		string(PrinterProcessingStateDetail_SeparationCutterResourceRemoved),
		string(PrinterProcessingStateDetail_SeparationCutterThermistorFailure),
		string(PrinterProcessingStateDetail_SeparationCutterTimingFailure),
		string(PrinterProcessingStateDetail_SeparationCutterTurnedOff),
		string(PrinterProcessingStateDetail_SeparationCutterTurnedOn),
		string(PrinterProcessingStateDetail_SeparationCutterUnderTemperature),
		string(PrinterProcessingStateDetail_SeparationCutterUnrecoverableFailure),
		string(PrinterProcessingStateDetail_SeparationCutterUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_SeparationCutterWarmingUp),
		string(PrinterProcessingStateDetail_SheetRotatorAdded),
		string(PrinterProcessingStateDetail_SheetRotatorAlmostEmpty),
		string(PrinterProcessingStateDetail_SheetRotatorAlmostFull),
		string(PrinterProcessingStateDetail_SheetRotatorAtLimit),
		string(PrinterProcessingStateDetail_SheetRotatorClosed),
		string(PrinterProcessingStateDetail_SheetRotatorConfigurationChange),
		string(PrinterProcessingStateDetail_SheetRotatorCoverClosed),
		string(PrinterProcessingStateDetail_SheetRotatorCoverOpen),
		string(PrinterProcessingStateDetail_SheetRotatorEmpty),
		string(PrinterProcessingStateDetail_SheetRotatorFull),
		string(PrinterProcessingStateDetail_SheetRotatorInterlockClosed),
		string(PrinterProcessingStateDetail_SheetRotatorInterlockOpen),
		string(PrinterProcessingStateDetail_SheetRotatorJam),
		string(PrinterProcessingStateDetail_SheetRotatorLifeAlmostOver),
		string(PrinterProcessingStateDetail_SheetRotatorLifeOver),
		string(PrinterProcessingStateDetail_SheetRotatorMemoryExhausted),
		string(PrinterProcessingStateDetail_SheetRotatorMissing),
		string(PrinterProcessingStateDetail_SheetRotatorMotorFailure),
		string(PrinterProcessingStateDetail_SheetRotatorNearLimit),
		string(PrinterProcessingStateDetail_SheetRotatorOffline),
		string(PrinterProcessingStateDetail_SheetRotatorOpened),
		string(PrinterProcessingStateDetail_SheetRotatorOverTemperature),
		string(PrinterProcessingStateDetail_SheetRotatorPowerSaver),
		string(PrinterProcessingStateDetail_SheetRotatorRecoverableFailure),
		string(PrinterProcessingStateDetail_SheetRotatorRecoverableStorage),
		string(PrinterProcessingStateDetail_SheetRotatorRemoved),
		string(PrinterProcessingStateDetail_SheetRotatorResourceAdded),
		string(PrinterProcessingStateDetail_SheetRotatorResourceRemoved),
		string(PrinterProcessingStateDetail_SheetRotatorThermistorFailure),
		string(PrinterProcessingStateDetail_SheetRotatorTimingFailure),
		string(PrinterProcessingStateDetail_SheetRotatorTurnedOff),
		string(PrinterProcessingStateDetail_SheetRotatorTurnedOn),
		string(PrinterProcessingStateDetail_SheetRotatorUnderTemperature),
		string(PrinterProcessingStateDetail_SheetRotatorUnrecoverableFailure),
		string(PrinterProcessingStateDetail_SheetRotatorUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_SheetRotatorWarmingUp),
		string(PrinterProcessingStateDetail_Shutdown),
		string(PrinterProcessingStateDetail_SlitterAdded),
		string(PrinterProcessingStateDetail_SlitterAlmostEmpty),
		string(PrinterProcessingStateDetail_SlitterAlmostFull),
		string(PrinterProcessingStateDetail_SlitterAtLimit),
		string(PrinterProcessingStateDetail_SlitterClosed),
		string(PrinterProcessingStateDetail_SlitterConfigurationChange),
		string(PrinterProcessingStateDetail_SlitterCoverClosed),
		string(PrinterProcessingStateDetail_SlitterCoverOpen),
		string(PrinterProcessingStateDetail_SlitterEmpty),
		string(PrinterProcessingStateDetail_SlitterFull),
		string(PrinterProcessingStateDetail_SlitterInterlockClosed),
		string(PrinterProcessingStateDetail_SlitterInterlockOpen),
		string(PrinterProcessingStateDetail_SlitterJam),
		string(PrinterProcessingStateDetail_SlitterLifeAlmostOver),
		string(PrinterProcessingStateDetail_SlitterLifeOver),
		string(PrinterProcessingStateDetail_SlitterMemoryExhausted),
		string(PrinterProcessingStateDetail_SlitterMissing),
		string(PrinterProcessingStateDetail_SlitterMotorFailure),
		string(PrinterProcessingStateDetail_SlitterNearLimit),
		string(PrinterProcessingStateDetail_SlitterOffline),
		string(PrinterProcessingStateDetail_SlitterOpened),
		string(PrinterProcessingStateDetail_SlitterOverTemperature),
		string(PrinterProcessingStateDetail_SlitterPowerSaver),
		string(PrinterProcessingStateDetail_SlitterRecoverableFailure),
		string(PrinterProcessingStateDetail_SlitterRecoverableStorage),
		string(PrinterProcessingStateDetail_SlitterRemoved),
		string(PrinterProcessingStateDetail_SlitterResourceAdded),
		string(PrinterProcessingStateDetail_SlitterResourceRemoved),
		string(PrinterProcessingStateDetail_SlitterThermistorFailure),
		string(PrinterProcessingStateDetail_SlitterTimingFailure),
		string(PrinterProcessingStateDetail_SlitterTurnedOff),
		string(PrinterProcessingStateDetail_SlitterTurnedOn),
		string(PrinterProcessingStateDetail_SlitterUnderTemperature),
		string(PrinterProcessingStateDetail_SlitterUnrecoverableFailure),
		string(PrinterProcessingStateDetail_SlitterUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_SlitterWarmingUp),
		string(PrinterProcessingStateDetail_SpoolAreaFull),
		string(PrinterProcessingStateDetail_StackerAdded),
		string(PrinterProcessingStateDetail_StackerAlmostEmpty),
		string(PrinterProcessingStateDetail_StackerAlmostFull),
		string(PrinterProcessingStateDetail_StackerAtLimit),
		string(PrinterProcessingStateDetail_StackerClosed),
		string(PrinterProcessingStateDetail_StackerConfigurationChange),
		string(PrinterProcessingStateDetail_StackerCoverClosed),
		string(PrinterProcessingStateDetail_StackerCoverOpen),
		string(PrinterProcessingStateDetail_StackerEmpty),
		string(PrinterProcessingStateDetail_StackerFull),
		string(PrinterProcessingStateDetail_StackerInterlockClosed),
		string(PrinterProcessingStateDetail_StackerInterlockOpen),
		string(PrinterProcessingStateDetail_StackerJam),
		string(PrinterProcessingStateDetail_StackerLifeAlmostOver),
		string(PrinterProcessingStateDetail_StackerLifeOver),
		string(PrinterProcessingStateDetail_StackerMemoryExhausted),
		string(PrinterProcessingStateDetail_StackerMissing),
		string(PrinterProcessingStateDetail_StackerMotorFailure),
		string(PrinterProcessingStateDetail_StackerNearLimit),
		string(PrinterProcessingStateDetail_StackerOffline),
		string(PrinterProcessingStateDetail_StackerOpened),
		string(PrinterProcessingStateDetail_StackerOverTemperature),
		string(PrinterProcessingStateDetail_StackerPowerSaver),
		string(PrinterProcessingStateDetail_StackerRecoverableFailure),
		string(PrinterProcessingStateDetail_StackerRecoverableStorage),
		string(PrinterProcessingStateDetail_StackerRemoved),
		string(PrinterProcessingStateDetail_StackerResourceAdded),
		string(PrinterProcessingStateDetail_StackerResourceRemoved),
		string(PrinterProcessingStateDetail_StackerThermistorFailure),
		string(PrinterProcessingStateDetail_StackerTimingFailure),
		string(PrinterProcessingStateDetail_StackerTurnedOff),
		string(PrinterProcessingStateDetail_StackerTurnedOn),
		string(PrinterProcessingStateDetail_StackerUnderTemperature),
		string(PrinterProcessingStateDetail_StackerUnrecoverableFailure),
		string(PrinterProcessingStateDetail_StackerUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_StackerWarmingUp),
		string(PrinterProcessingStateDetail_Standby),
		string(PrinterProcessingStateDetail_StaplerAdded),
		string(PrinterProcessingStateDetail_StaplerAlmostEmpty),
		string(PrinterProcessingStateDetail_StaplerAlmostFull),
		string(PrinterProcessingStateDetail_StaplerAtLimit),
		string(PrinterProcessingStateDetail_StaplerClosed),
		string(PrinterProcessingStateDetail_StaplerConfigurationChange),
		string(PrinterProcessingStateDetail_StaplerCoverClosed),
		string(PrinterProcessingStateDetail_StaplerCoverOpen),
		string(PrinterProcessingStateDetail_StaplerEmpty),
		string(PrinterProcessingStateDetail_StaplerFull),
		string(PrinterProcessingStateDetail_StaplerInterlockClosed),
		string(PrinterProcessingStateDetail_StaplerInterlockOpen),
		string(PrinterProcessingStateDetail_StaplerJam),
		string(PrinterProcessingStateDetail_StaplerLifeAlmostOver),
		string(PrinterProcessingStateDetail_StaplerLifeOver),
		string(PrinterProcessingStateDetail_StaplerMemoryExhausted),
		string(PrinterProcessingStateDetail_StaplerMissing),
		string(PrinterProcessingStateDetail_StaplerMotorFailure),
		string(PrinterProcessingStateDetail_StaplerNearLimit),
		string(PrinterProcessingStateDetail_StaplerOffline),
		string(PrinterProcessingStateDetail_StaplerOpened),
		string(PrinterProcessingStateDetail_StaplerOverTemperature),
		string(PrinterProcessingStateDetail_StaplerPowerSaver),
		string(PrinterProcessingStateDetail_StaplerRecoverableFailure),
		string(PrinterProcessingStateDetail_StaplerRecoverableStorage),
		string(PrinterProcessingStateDetail_StaplerRemoved),
		string(PrinterProcessingStateDetail_StaplerResourceAdded),
		string(PrinterProcessingStateDetail_StaplerResourceRemoved),
		string(PrinterProcessingStateDetail_StaplerThermistorFailure),
		string(PrinterProcessingStateDetail_StaplerTimingFailure),
		string(PrinterProcessingStateDetail_StaplerTurnedOff),
		string(PrinterProcessingStateDetail_StaplerTurnedOn),
		string(PrinterProcessingStateDetail_StaplerUnderTemperature),
		string(PrinterProcessingStateDetail_StaplerUnrecoverableFailure),
		string(PrinterProcessingStateDetail_StaplerUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_StaplerWarmingUp),
		string(PrinterProcessingStateDetail_StitcherAdded),
		string(PrinterProcessingStateDetail_StitcherAlmostEmpty),
		string(PrinterProcessingStateDetail_StitcherAlmostFull),
		string(PrinterProcessingStateDetail_StitcherAtLimit),
		string(PrinterProcessingStateDetail_StitcherClosed),
		string(PrinterProcessingStateDetail_StitcherConfigurationChange),
		string(PrinterProcessingStateDetail_StitcherCoverClosed),
		string(PrinterProcessingStateDetail_StitcherCoverOpen),
		string(PrinterProcessingStateDetail_StitcherEmpty),
		string(PrinterProcessingStateDetail_StitcherFull),
		string(PrinterProcessingStateDetail_StitcherInterlockClosed),
		string(PrinterProcessingStateDetail_StitcherInterlockOpen),
		string(PrinterProcessingStateDetail_StitcherJam),
		string(PrinterProcessingStateDetail_StitcherLifeAlmostOver),
		string(PrinterProcessingStateDetail_StitcherLifeOver),
		string(PrinterProcessingStateDetail_StitcherMemoryExhausted),
		string(PrinterProcessingStateDetail_StitcherMissing),
		string(PrinterProcessingStateDetail_StitcherMotorFailure),
		string(PrinterProcessingStateDetail_StitcherNearLimit),
		string(PrinterProcessingStateDetail_StitcherOffline),
		string(PrinterProcessingStateDetail_StitcherOpened),
		string(PrinterProcessingStateDetail_StitcherOverTemperature),
		string(PrinterProcessingStateDetail_StitcherPowerSaver),
		string(PrinterProcessingStateDetail_StitcherRecoverableFailure),
		string(PrinterProcessingStateDetail_StitcherRecoverableStorage),
		string(PrinterProcessingStateDetail_StitcherRemoved),
		string(PrinterProcessingStateDetail_StitcherResourceAdded),
		string(PrinterProcessingStateDetail_StitcherResourceRemoved),
		string(PrinterProcessingStateDetail_StitcherThermistorFailure),
		string(PrinterProcessingStateDetail_StitcherTimingFailure),
		string(PrinterProcessingStateDetail_StitcherTurnedOff),
		string(PrinterProcessingStateDetail_StitcherTurnedOn),
		string(PrinterProcessingStateDetail_StitcherUnderTemperature),
		string(PrinterProcessingStateDetail_StitcherUnrecoverableFailure),
		string(PrinterProcessingStateDetail_StitcherUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_StitcherWarmingUp),
		string(PrinterProcessingStateDetail_StoppedPartially),
		string(PrinterProcessingStateDetail_Stopping),
		string(PrinterProcessingStateDetail_SubunitAdded),
		string(PrinterProcessingStateDetail_SubunitAlmostEmpty),
		string(PrinterProcessingStateDetail_SubunitAlmostFull),
		string(PrinterProcessingStateDetail_SubunitAtLimit),
		string(PrinterProcessingStateDetail_SubunitClosed),
		string(PrinterProcessingStateDetail_SubunitCoolingDown),
		string(PrinterProcessingStateDetail_SubunitEmpty),
		string(PrinterProcessingStateDetail_SubunitFull),
		string(PrinterProcessingStateDetail_SubunitLifeAlmostOver),
		string(PrinterProcessingStateDetail_SubunitLifeOver),
		string(PrinterProcessingStateDetail_SubunitMemoryExhausted),
		string(PrinterProcessingStateDetail_SubunitMissing),
		string(PrinterProcessingStateDetail_SubunitMotorFailure),
		string(PrinterProcessingStateDetail_SubunitNearLimit),
		string(PrinterProcessingStateDetail_SubunitOffline),
		string(PrinterProcessingStateDetail_SubunitOpened),
		string(PrinterProcessingStateDetail_SubunitOverTemperature),
		string(PrinterProcessingStateDetail_SubunitPowerSaver),
		string(PrinterProcessingStateDetail_SubunitRecoverableFailure),
		string(PrinterProcessingStateDetail_SubunitRecoverableStorage),
		string(PrinterProcessingStateDetail_SubunitRemoved),
		string(PrinterProcessingStateDetail_SubunitResourceAdded),
		string(PrinterProcessingStateDetail_SubunitResourceRemoved),
		string(PrinterProcessingStateDetail_SubunitThermistorFailure),
		string(PrinterProcessingStateDetail_SubunitTimingFailure),
		string(PrinterProcessingStateDetail_SubunitTurnedOff),
		string(PrinterProcessingStateDetail_SubunitTurnedOn),
		string(PrinterProcessingStateDetail_SubunitUnderTemperature),
		string(PrinterProcessingStateDetail_SubunitUnrecoverableFailure),
		string(PrinterProcessingStateDetail_SubunitUnrecoverableStorage),
		string(PrinterProcessingStateDetail_SubunitWarmingUp),
		string(PrinterProcessingStateDetail_Suspend),
		string(PrinterProcessingStateDetail_Testing),
		string(PrinterProcessingStateDetail_TimedOut),
		string(PrinterProcessingStateDetail_TonerEmpty),
		string(PrinterProcessingStateDetail_TonerLow),
		string(PrinterProcessingStateDetail_TrimmerAdded),
		string(PrinterProcessingStateDetail_TrimmerAlmostEmpty),
		string(PrinterProcessingStateDetail_TrimmerAlmostFull),
		string(PrinterProcessingStateDetail_TrimmerAtLimit),
		string(PrinterProcessingStateDetail_TrimmerClosed),
		string(PrinterProcessingStateDetail_TrimmerConfigurationChange),
		string(PrinterProcessingStateDetail_TrimmerCoverClosed),
		string(PrinterProcessingStateDetail_TrimmerCoverOpen),
		string(PrinterProcessingStateDetail_TrimmerEmpty),
		string(PrinterProcessingStateDetail_TrimmerFull),
		string(PrinterProcessingStateDetail_TrimmerInterlockClosed),
		string(PrinterProcessingStateDetail_TrimmerInterlockOpen),
		string(PrinterProcessingStateDetail_TrimmerJam),
		string(PrinterProcessingStateDetail_TrimmerLifeAlmostOver),
		string(PrinterProcessingStateDetail_TrimmerLifeOver),
		string(PrinterProcessingStateDetail_TrimmerMemoryExhausted),
		string(PrinterProcessingStateDetail_TrimmerMissing),
		string(PrinterProcessingStateDetail_TrimmerMotorFailure),
		string(PrinterProcessingStateDetail_TrimmerNearLimit),
		string(PrinterProcessingStateDetail_TrimmerOffline),
		string(PrinterProcessingStateDetail_TrimmerOpened),
		string(PrinterProcessingStateDetail_TrimmerOverTemperature),
		string(PrinterProcessingStateDetail_TrimmerPowerSaver),
		string(PrinterProcessingStateDetail_TrimmerRecoverableFailure),
		string(PrinterProcessingStateDetail_TrimmerRecoverableStorage),
		string(PrinterProcessingStateDetail_TrimmerRemoved),
		string(PrinterProcessingStateDetail_TrimmerResourceAdded),
		string(PrinterProcessingStateDetail_TrimmerResourceRemoved),
		string(PrinterProcessingStateDetail_TrimmerThermistorFailure),
		string(PrinterProcessingStateDetail_TrimmerTimingFailure),
		string(PrinterProcessingStateDetail_TrimmerTurnedOff),
		string(PrinterProcessingStateDetail_TrimmerTurnedOn),
		string(PrinterProcessingStateDetail_TrimmerUnderTemperature),
		string(PrinterProcessingStateDetail_TrimmerUnrecoverableFailure),
		string(PrinterProcessingStateDetail_TrimmerUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_TrimmerWarmingUp),
		string(PrinterProcessingStateDetail_Unknown),
		string(PrinterProcessingStateDetail_WrapperAdded),
		string(PrinterProcessingStateDetail_WrapperAlmostEmpty),
		string(PrinterProcessingStateDetail_WrapperAlmostFull),
		string(PrinterProcessingStateDetail_WrapperAtLimit),
		string(PrinterProcessingStateDetail_WrapperClosed),
		string(PrinterProcessingStateDetail_WrapperConfigurationChange),
		string(PrinterProcessingStateDetail_WrapperCoverClosed),
		string(PrinterProcessingStateDetail_WrapperCoverOpen),
		string(PrinterProcessingStateDetail_WrapperEmpty),
		string(PrinterProcessingStateDetail_WrapperFull),
		string(PrinterProcessingStateDetail_WrapperInterlockClosed),
		string(PrinterProcessingStateDetail_WrapperInterlockOpen),
		string(PrinterProcessingStateDetail_WrapperJam),
		string(PrinterProcessingStateDetail_WrapperLifeAlmostOver),
		string(PrinterProcessingStateDetail_WrapperLifeOver),
		string(PrinterProcessingStateDetail_WrapperMemoryExhausted),
		string(PrinterProcessingStateDetail_WrapperMissing),
		string(PrinterProcessingStateDetail_WrapperMotorFailure),
		string(PrinterProcessingStateDetail_WrapperNearLimit),
		string(PrinterProcessingStateDetail_WrapperOffline),
		string(PrinterProcessingStateDetail_WrapperOpened),
		string(PrinterProcessingStateDetail_WrapperOverTemperature),
		string(PrinterProcessingStateDetail_WrapperPowerSaver),
		string(PrinterProcessingStateDetail_WrapperRecoverableFailure),
		string(PrinterProcessingStateDetail_WrapperRecoverableStorage),
		string(PrinterProcessingStateDetail_WrapperRemoved),
		string(PrinterProcessingStateDetail_WrapperResourceAdded),
		string(PrinterProcessingStateDetail_WrapperResourceRemoved),
		string(PrinterProcessingStateDetail_WrapperThermistorFailure),
		string(PrinterProcessingStateDetail_WrapperTimingFailure),
		string(PrinterProcessingStateDetail_WrapperTurnedOff),
		string(PrinterProcessingStateDetail_WrapperTurnedOn),
		string(PrinterProcessingStateDetail_WrapperUnderTemperature),
		string(PrinterProcessingStateDetail_WrapperUnrecoverableFailure),
		string(PrinterProcessingStateDetail_WrapperUnrecoverableStorageError),
		string(PrinterProcessingStateDetail_WrapperWarmingUp),
	}
}

func (s *PrinterProcessingStateDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterProcessingStateDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterProcessingStateDetail(input string) (*PrinterProcessingStateDetail, error) {
	vals := map[string]PrinterProcessingStateDetail{
		"alertremovalofbinarychangeentry":           PrinterProcessingStateDetail_AlertRemovalOfBinaryChangeEntry,
		"banderadded":                               PrinterProcessingStateDetail_BanderAdded,
		"banderalmostempty":                         PrinterProcessingStateDetail_BanderAlmostEmpty,
		"banderalmostfull":                          PrinterProcessingStateDetail_BanderAlmostFull,
		"banderatlimit":                             PrinterProcessingStateDetail_BanderAtLimit,
		"banderclosed":                              PrinterProcessingStateDetail_BanderClosed,
		"banderconfigurationchange":                 PrinterProcessingStateDetail_BanderConfigurationChange,
		"bandercoverclosed":                         PrinterProcessingStateDetail_BanderCoverClosed,
		"bandercoveropen":                           PrinterProcessingStateDetail_BanderCoverOpen,
		"banderempty":                               PrinterProcessingStateDetail_BanderEmpty,
		"banderfull":                                PrinterProcessingStateDetail_BanderFull,
		"banderinterlockclosed":                     PrinterProcessingStateDetail_BanderInterlockClosed,
		"banderinterlockopen":                       PrinterProcessingStateDetail_BanderInterlockOpen,
		"banderjam":                                 PrinterProcessingStateDetail_BanderJam,
		"banderlifealmostover":                      PrinterProcessingStateDetail_BanderLifeAlmostOver,
		"banderlifeover":                            PrinterProcessingStateDetail_BanderLifeOver,
		"bandermemoryexhausted":                     PrinterProcessingStateDetail_BanderMemoryExhausted,
		"bandermissing":                             PrinterProcessingStateDetail_BanderMissing,
		"bandermotorfailure":                        PrinterProcessingStateDetail_BanderMotorFailure,
		"bandernearlimit":                           PrinterProcessingStateDetail_BanderNearLimit,
		"banderoffline":                             PrinterProcessingStateDetail_BanderOffline,
		"banderopened":                              PrinterProcessingStateDetail_BanderOpened,
		"banderovertemperature":                     PrinterProcessingStateDetail_BanderOverTemperature,
		"banderpowersaver":                          PrinterProcessingStateDetail_BanderPowerSaver,
		"banderrecoverablefailure":                  PrinterProcessingStateDetail_BanderRecoverableFailure,
		"banderrecoverablestorage":                  PrinterProcessingStateDetail_BanderRecoverableStorage,
		"banderremoved":                             PrinterProcessingStateDetail_BanderRemoved,
		"banderresourceadded":                       PrinterProcessingStateDetail_BanderResourceAdded,
		"banderresourceremoved":                     PrinterProcessingStateDetail_BanderResourceRemoved,
		"banderthermistorfailure":                   PrinterProcessingStateDetail_BanderThermistorFailure,
		"bandertimingfailure":                       PrinterProcessingStateDetail_BanderTimingFailure,
		"banderturnedoff":                           PrinterProcessingStateDetail_BanderTurnedOff,
		"banderturnedon":                            PrinterProcessingStateDetail_BanderTurnedOn,
		"banderundertemperature":                    PrinterProcessingStateDetail_BanderUnderTemperature,
		"banderunrecoverablefailure":                PrinterProcessingStateDetail_BanderUnrecoverableFailure,
		"banderunrecoverablestorageerror":           PrinterProcessingStateDetail_BanderUnrecoverableStorageError,
		"banderwarmingup":                           PrinterProcessingStateDetail_BanderWarmingUp,
		"binderadded":                               PrinterProcessingStateDetail_BinderAdded,
		"binderalmostempty":                         PrinterProcessingStateDetail_BinderAlmostEmpty,
		"binderalmostfull":                          PrinterProcessingStateDetail_BinderAlmostFull,
		"binderatlimit":                             PrinterProcessingStateDetail_BinderAtLimit,
		"binderclosed":                              PrinterProcessingStateDetail_BinderClosed,
		"binderconfigurationchange":                 PrinterProcessingStateDetail_BinderConfigurationChange,
		"bindercoverclosed":                         PrinterProcessingStateDetail_BinderCoverClosed,
		"bindercoveropen":                           PrinterProcessingStateDetail_BinderCoverOpen,
		"binderempty":                               PrinterProcessingStateDetail_BinderEmpty,
		"binderfull":                                PrinterProcessingStateDetail_BinderFull,
		"binderinterlockclosed":                     PrinterProcessingStateDetail_BinderInterlockClosed,
		"binderinterlockopen":                       PrinterProcessingStateDetail_BinderInterlockOpen,
		"binderjam":                                 PrinterProcessingStateDetail_BinderJam,
		"binderlifealmostover":                      PrinterProcessingStateDetail_BinderLifeAlmostOver,
		"binderlifeover":                            PrinterProcessingStateDetail_BinderLifeOver,
		"bindermemoryexhausted":                     PrinterProcessingStateDetail_BinderMemoryExhausted,
		"bindermissing":                             PrinterProcessingStateDetail_BinderMissing,
		"bindermotorfailure":                        PrinterProcessingStateDetail_BinderMotorFailure,
		"bindernearlimit":                           PrinterProcessingStateDetail_BinderNearLimit,
		"binderoffline":                             PrinterProcessingStateDetail_BinderOffline,
		"binderopened":                              PrinterProcessingStateDetail_BinderOpened,
		"binderovertemperature":                     PrinterProcessingStateDetail_BinderOverTemperature,
		"binderpowersaver":                          PrinterProcessingStateDetail_BinderPowerSaver,
		"binderrecoverablefailure":                  PrinterProcessingStateDetail_BinderRecoverableFailure,
		"binderrecoverablestorage":                  PrinterProcessingStateDetail_BinderRecoverableStorage,
		"binderremoved":                             PrinterProcessingStateDetail_BinderRemoved,
		"binderresourceadded":                       PrinterProcessingStateDetail_BinderResourceAdded,
		"binderresourceremoved":                     PrinterProcessingStateDetail_BinderResourceRemoved,
		"binderthermistorfailure":                   PrinterProcessingStateDetail_BinderThermistorFailure,
		"bindertimingfailure":                       PrinterProcessingStateDetail_BinderTimingFailure,
		"binderturnedoff":                           PrinterProcessingStateDetail_BinderTurnedOff,
		"binderturnedon":                            PrinterProcessingStateDetail_BinderTurnedOn,
		"binderundertemperature":                    PrinterProcessingStateDetail_BinderUnderTemperature,
		"binderunrecoverablefailure":                PrinterProcessingStateDetail_BinderUnrecoverableFailure,
		"binderunrecoverablestorageerror":           PrinterProcessingStateDetail_BinderUnrecoverableStorageError,
		"binderwarmingup":                           PrinterProcessingStateDetail_BinderWarmingUp,
		"camerafailure":                             PrinterProcessingStateDetail_CameraFailure,
		"chambercooling":                            PrinterProcessingStateDetail_ChamberCooling,
		"chamberfailure":                            PrinterProcessingStateDetail_ChamberFailure,
		"chamberheating":                            PrinterProcessingStateDetail_ChamberHeating,
		"chambertemperaturehigh":                    PrinterProcessingStateDetail_ChamberTemperatureHigh,
		"chambertemperaturelow":                     PrinterProcessingStateDetail_ChamberTemperatureLow,
		"cleanerlifealmostover":                     PrinterProcessingStateDetail_CleanerLifeAlmostOver,
		"cleanerlifeover":                           PrinterProcessingStateDetail_CleanerLifeOver,
		"configurationchange":                       PrinterProcessingStateDetail_ConfigurationChange,
		"connectingtodevice":                        PrinterProcessingStateDetail_ConnectingToDevice,
		"coveropen":                                 PrinterProcessingStateDetail_CoverOpen,
		"deactivated":                               PrinterProcessingStateDetail_Deactivated,
		"deleted":                                   PrinterProcessingStateDetail_Deleted,
		"developerempty":                            PrinterProcessingStateDetail_DeveloperEmpty,
		"developerlow":                              PrinterProcessingStateDetail_DeveloperLow,
		"diecutteradded":                            PrinterProcessingStateDetail_DieCutterAdded,
		"diecutteralmostempty":                      PrinterProcessingStateDetail_DieCutterAlmostEmpty,
		"diecutteralmostfull":                       PrinterProcessingStateDetail_DieCutterAlmostFull,
		"diecutteratlimit":                          PrinterProcessingStateDetail_DieCutterAtLimit,
		"diecutterclosed":                           PrinterProcessingStateDetail_DieCutterClosed,
		"diecutterconfigurationchange":              PrinterProcessingStateDetail_DieCutterConfigurationChange,
		"diecuttercoverclosed":                      PrinterProcessingStateDetail_DieCutterCoverClosed,
		"diecuttercoveropen":                        PrinterProcessingStateDetail_DieCutterCoverOpen,
		"diecutterempty":                            PrinterProcessingStateDetail_DieCutterEmpty,
		"diecutterfull":                             PrinterProcessingStateDetail_DieCutterFull,
		"diecutterinterlockclosed":                  PrinterProcessingStateDetail_DieCutterInterlockClosed,
		"diecutterinterlockopen":                    PrinterProcessingStateDetail_DieCutterInterlockOpen,
		"diecutterjam":                              PrinterProcessingStateDetail_DieCutterJam,
		"diecutterlifealmostover":                   PrinterProcessingStateDetail_DieCutterLifeAlmostOver,
		"diecutterlifeover":                         PrinterProcessingStateDetail_DieCutterLifeOver,
		"diecuttermemoryexhausted":                  PrinterProcessingStateDetail_DieCutterMemoryExhausted,
		"diecuttermissing":                          PrinterProcessingStateDetail_DieCutterMissing,
		"diecuttermotorfailure":                     PrinterProcessingStateDetail_DieCutterMotorFailure,
		"diecutternearlimit":                        PrinterProcessingStateDetail_DieCutterNearLimit,
		"diecutteroffline":                          PrinterProcessingStateDetail_DieCutterOffline,
		"diecutteropened":                           PrinterProcessingStateDetail_DieCutterOpened,
		"diecutterovertemperature":                  PrinterProcessingStateDetail_DieCutterOverTemperature,
		"diecutterpowersaver":                       PrinterProcessingStateDetail_DieCutterPowerSaver,
		"diecutterrecoverablefailure":               PrinterProcessingStateDetail_DieCutterRecoverableFailure,
		"diecutterrecoverablestorage":               PrinterProcessingStateDetail_DieCutterRecoverableStorage,
		"diecutterremoved":                          PrinterProcessingStateDetail_DieCutterRemoved,
		"diecutterresourceadded":                    PrinterProcessingStateDetail_DieCutterResourceAdded,
		"diecutterresourceremoved":                  PrinterProcessingStateDetail_DieCutterResourceRemoved,
		"diecutterthermistorfailure":                PrinterProcessingStateDetail_DieCutterThermistorFailure,
		"diecuttertimingfailure":                    PrinterProcessingStateDetail_DieCutterTimingFailure,
		"diecutterturnedoff":                        PrinterProcessingStateDetail_DieCutterTurnedOff,
		"diecutterturnedon":                         PrinterProcessingStateDetail_DieCutterTurnedOn,
		"diecutterundertemperature":                 PrinterProcessingStateDetail_DieCutterUnderTemperature,
		"diecutterunrecoverablefailure":             PrinterProcessingStateDetail_DieCutterUnrecoverableFailure,
		"diecutterunrecoverablestorageerror":        PrinterProcessingStateDetail_DieCutterUnrecoverableStorageError,
		"diecutterwarmingup":                        PrinterProcessingStateDetail_DieCutterWarmingUp,
		"dooropen":                                  PrinterProcessingStateDetail_DoorOpen,
		"extrudercooling":                           PrinterProcessingStateDetail_ExtruderCooling,
		"extruderfailure":                           PrinterProcessingStateDetail_ExtruderFailure,
		"extruderheating":                           PrinterProcessingStateDetail_ExtruderHeating,
		"extruderjam":                               PrinterProcessingStateDetail_ExtruderJam,
		"extrudertemperaturehigh":                   PrinterProcessingStateDetail_ExtruderTemperatureHigh,
		"extrudertemperaturelow":                    PrinterProcessingStateDetail_ExtruderTemperatureLow,
		"fanfailure":                                PrinterProcessingStateDetail_FanFailure,
		"faxmodemlifealmostover":                    PrinterProcessingStateDetail_FaxModemLifeAlmostOver,
		"faxmodemlifeover":                          PrinterProcessingStateDetail_FaxModemLifeOver,
		"faxmodemmissing":                           PrinterProcessingStateDetail_FaxModemMissing,
		"faxmodemturnedoff":                         PrinterProcessingStateDetail_FaxModemTurnedOff,
		"faxmodemturnedon":                          PrinterProcessingStateDetail_FaxModemTurnedOn,
		"folderadded":                               PrinterProcessingStateDetail_FolderAdded,
		"folderalmostempty":                         PrinterProcessingStateDetail_FolderAlmostEmpty,
		"folderalmostfull":                          PrinterProcessingStateDetail_FolderAlmostFull,
		"folderatlimit":                             PrinterProcessingStateDetail_FolderAtLimit,
		"folderclosed":                              PrinterProcessingStateDetail_FolderClosed,
		"folderconfigurationchange":                 PrinterProcessingStateDetail_FolderConfigurationChange,
		"foldercoverclosed":                         PrinterProcessingStateDetail_FolderCoverClosed,
		"foldercoveropen":                           PrinterProcessingStateDetail_FolderCoverOpen,
		"folderempty":                               PrinterProcessingStateDetail_FolderEmpty,
		"folderfull":                                PrinterProcessingStateDetail_FolderFull,
		"folderinterlockclosed":                     PrinterProcessingStateDetail_FolderInterlockClosed,
		"folderinterlockopen":                       PrinterProcessingStateDetail_FolderInterlockOpen,
		"folderjam":                                 PrinterProcessingStateDetail_FolderJam,
		"folderlifealmostover":                      PrinterProcessingStateDetail_FolderLifeAlmostOver,
		"folderlifeover":                            PrinterProcessingStateDetail_FolderLifeOver,
		"foldermemoryexhausted":                     PrinterProcessingStateDetail_FolderMemoryExhausted,
		"foldermissing":                             PrinterProcessingStateDetail_FolderMissing,
		"foldermotorfailure":                        PrinterProcessingStateDetail_FolderMotorFailure,
		"foldernearlimit":                           PrinterProcessingStateDetail_FolderNearLimit,
		"folderoffline":                             PrinterProcessingStateDetail_FolderOffline,
		"folderopened":                              PrinterProcessingStateDetail_FolderOpened,
		"folderovertemperature":                     PrinterProcessingStateDetail_FolderOverTemperature,
		"folderpowersaver":                          PrinterProcessingStateDetail_FolderPowerSaver,
		"folderrecoverablefailure":                  PrinterProcessingStateDetail_FolderRecoverableFailure,
		"folderrecoverablestorage":                  PrinterProcessingStateDetail_FolderRecoverableStorage,
		"folderremoved":                             PrinterProcessingStateDetail_FolderRemoved,
		"folderresourceadded":                       PrinterProcessingStateDetail_FolderResourceAdded,
		"folderresourceremoved":                     PrinterProcessingStateDetail_FolderResourceRemoved,
		"folderthermistorfailure":                   PrinterProcessingStateDetail_FolderThermistorFailure,
		"foldertimingfailure":                       PrinterProcessingStateDetail_FolderTimingFailure,
		"folderturnedoff":                           PrinterProcessingStateDetail_FolderTurnedOff,
		"folderturnedon":                            PrinterProcessingStateDetail_FolderTurnedOn,
		"folderundertemperature":                    PrinterProcessingStateDetail_FolderUnderTemperature,
		"folderunrecoverablefailure":                PrinterProcessingStateDetail_FolderUnrecoverableFailure,
		"folderunrecoverablestorageerror":           PrinterProcessingStateDetail_FolderUnrecoverableStorageError,
		"folderwarmingup":                           PrinterProcessingStateDetail_FolderWarmingUp,
		"fuserovertemp":                             PrinterProcessingStateDetail_FuserOverTemp,
		"fuserundertemp":                            PrinterProcessingStateDetail_FuserUnderTemp,
		"hibernate":                                 PrinterProcessingStateDetail_Hibernate,
		"holdnewjobs":                               PrinterProcessingStateDetail_HoldNewJobs,
		"identifyprinterrequested":                  PrinterProcessingStateDetail_IdentifyPrinterRequested,
		"imprinteradded":                            PrinterProcessingStateDetail_ImprinterAdded,
		"imprinteralmostempty":                      PrinterProcessingStateDetail_ImprinterAlmostEmpty,
		"imprinteralmostfull":                       PrinterProcessingStateDetail_ImprinterAlmostFull,
		"imprinteratlimit":                          PrinterProcessingStateDetail_ImprinterAtLimit,
		"imprinterclosed":                           PrinterProcessingStateDetail_ImprinterClosed,
		"imprinterconfigurationchange":              PrinterProcessingStateDetail_ImprinterConfigurationChange,
		"imprintercoverclosed":                      PrinterProcessingStateDetail_ImprinterCoverClosed,
		"imprintercoveropen":                        PrinterProcessingStateDetail_ImprinterCoverOpen,
		"imprinterempty":                            PrinterProcessingStateDetail_ImprinterEmpty,
		"imprinterfull":                             PrinterProcessingStateDetail_ImprinterFull,
		"imprinterinterlockclosed":                  PrinterProcessingStateDetail_ImprinterInterlockClosed,
		"imprinterinterlockopen":                    PrinterProcessingStateDetail_ImprinterInterlockOpen,
		"imprinterjam":                              PrinterProcessingStateDetail_ImprinterJam,
		"imprinterlifealmostover":                   PrinterProcessingStateDetail_ImprinterLifeAlmostOver,
		"imprinterlifeover":                         PrinterProcessingStateDetail_ImprinterLifeOver,
		"imprintermemoryexhausted":                  PrinterProcessingStateDetail_ImprinterMemoryExhausted,
		"imprintermissing":                          PrinterProcessingStateDetail_ImprinterMissing,
		"imprintermotorfailure":                     PrinterProcessingStateDetail_ImprinterMotorFailure,
		"imprinternearlimit":                        PrinterProcessingStateDetail_ImprinterNearLimit,
		"imprinteroffline":                          PrinterProcessingStateDetail_ImprinterOffline,
		"imprinteropened":                           PrinterProcessingStateDetail_ImprinterOpened,
		"imprinterovertemperature":                  PrinterProcessingStateDetail_ImprinterOverTemperature,
		"imprinterpowersaver":                       PrinterProcessingStateDetail_ImprinterPowerSaver,
		"imprinterrecoverablefailure":               PrinterProcessingStateDetail_ImprinterRecoverableFailure,
		"imprinterrecoverablestorage":               PrinterProcessingStateDetail_ImprinterRecoverableStorage,
		"imprinterremoved":                          PrinterProcessingStateDetail_ImprinterRemoved,
		"imprinterresourceadded":                    PrinterProcessingStateDetail_ImprinterResourceAdded,
		"imprinterresourceremoved":                  PrinterProcessingStateDetail_ImprinterResourceRemoved,
		"imprinterthermistorfailure":                PrinterProcessingStateDetail_ImprinterThermistorFailure,
		"imprintertimingfailure":                    PrinterProcessingStateDetail_ImprinterTimingFailure,
		"imprinterturnedoff":                        PrinterProcessingStateDetail_ImprinterTurnedOff,
		"imprinterturnedon":                         PrinterProcessingStateDetail_ImprinterTurnedOn,
		"imprinterundertemperature":                 PrinterProcessingStateDetail_ImprinterUnderTemperature,
		"imprinterunrecoverablefailure":             PrinterProcessingStateDetail_ImprinterUnrecoverableFailure,
		"imprinterunrecoverablestorageerror":        PrinterProcessingStateDetail_ImprinterUnrecoverableStorageError,
		"imprinterwarmingup":                        PrinterProcessingStateDetail_ImprinterWarmingUp,
		"inputcannotfeedsizeselected":               PrinterProcessingStateDetail_InputCannotFeedSizeSelected,
		"inputmanualinputrequest":                   PrinterProcessingStateDetail_InputManualInputRequest,
		"inputmediacolorchange":                     PrinterProcessingStateDetail_InputMediaColorChange,
		"inputmediaformpartschange":                 PrinterProcessingStateDetail_InputMediaFormPartsChange,
		"inputmediasizechange":                      PrinterProcessingStateDetail_InputMediaSizeChange,
		"inputmediatrayfailure":                     PrinterProcessingStateDetail_InputMediaTrayFailure,
		"inputmediatrayfeederror":                   PrinterProcessingStateDetail_InputMediaTrayFeedError,
		"inputmediatrayjam":                         PrinterProcessingStateDetail_InputMediaTrayJam,
		"inputmediatypechange":                      PrinterProcessingStateDetail_InputMediaTypeChange,
		"inputmediaweightchange":                    PrinterProcessingStateDetail_InputMediaWeightChange,
		"inputpickrollerfailure":                    PrinterProcessingStateDetail_InputPickRollerFailure,
		"inputpickrollerlifeover":                   PrinterProcessingStateDetail_InputPickRollerLifeOver,
		"inputpickrollerlifewarn":                   PrinterProcessingStateDetail_InputPickRollerLifeWarn,
		"inputpickrollermissing":                    PrinterProcessingStateDetail_InputPickRollerMissing,
		"inputtrayelevationfailure":                 PrinterProcessingStateDetail_InputTrayElevationFailure,
		"inputtraymissing":                          PrinterProcessingStateDetail_InputTrayMissing,
		"inputtraypositionfailure":                  PrinterProcessingStateDetail_InputTrayPositionFailure,
		"inserteradded":                             PrinterProcessingStateDetail_InserterAdded,
		"inserteralmostempty":                       PrinterProcessingStateDetail_InserterAlmostEmpty,
		"inserteralmostfull":                        PrinterProcessingStateDetail_InserterAlmostFull,
		"inserteratlimit":                           PrinterProcessingStateDetail_InserterAtLimit,
		"inserterclosed":                            PrinterProcessingStateDetail_InserterClosed,
		"inserterconfigurationchange":               PrinterProcessingStateDetail_InserterConfigurationChange,
		"insertercoverclosed":                       PrinterProcessingStateDetail_InserterCoverClosed,
		"insertercoveropen":                         PrinterProcessingStateDetail_InserterCoverOpen,
		"inserterempty":                             PrinterProcessingStateDetail_InserterEmpty,
		"inserterfull":                              PrinterProcessingStateDetail_InserterFull,
		"inserterinterlockclosed":                   PrinterProcessingStateDetail_InserterInterlockClosed,
		"inserterinterlockopen":                     PrinterProcessingStateDetail_InserterInterlockOpen,
		"inserterjam":                               PrinterProcessingStateDetail_InserterJam,
		"inserterlifealmostover":                    PrinterProcessingStateDetail_InserterLifeAlmostOver,
		"inserterlifeover":                          PrinterProcessingStateDetail_InserterLifeOver,
		"insertermemoryexhausted":                   PrinterProcessingStateDetail_InserterMemoryExhausted,
		"insertermissing":                           PrinterProcessingStateDetail_InserterMissing,
		"insertermotorfailure":                      PrinterProcessingStateDetail_InserterMotorFailure,
		"inserternearlimit":                         PrinterProcessingStateDetail_InserterNearLimit,
		"inserteroffline":                           PrinterProcessingStateDetail_InserterOffline,
		"inserteropened":                            PrinterProcessingStateDetail_InserterOpened,
		"inserterovertemperature":                   PrinterProcessingStateDetail_InserterOverTemperature,
		"inserterpowersaver":                        PrinterProcessingStateDetail_InserterPowerSaver,
		"inserterrecoverablefailure":                PrinterProcessingStateDetail_InserterRecoverableFailure,
		"inserterrecoverablestorage":                PrinterProcessingStateDetail_InserterRecoverableStorage,
		"inserterremoved":                           PrinterProcessingStateDetail_InserterRemoved,
		"inserterresourceadded":                     PrinterProcessingStateDetail_InserterResourceAdded,
		"inserterresourceremoved":                   PrinterProcessingStateDetail_InserterResourceRemoved,
		"inserterthermistorfailure":                 PrinterProcessingStateDetail_InserterThermistorFailure,
		"insertertimingfailure":                     PrinterProcessingStateDetail_InserterTimingFailure,
		"inserterturnedoff":                         PrinterProcessingStateDetail_InserterTurnedOff,
		"inserterturnedon":                          PrinterProcessingStateDetail_InserterTurnedOn,
		"inserterundertemperature":                  PrinterProcessingStateDetail_InserterUnderTemperature,
		"inserterunrecoverablefailure":              PrinterProcessingStateDetail_InserterUnrecoverableFailure,
		"inserterunrecoverablestorageerror":         PrinterProcessingStateDetail_InserterUnrecoverableStorageError,
		"inserterwarmingup":                         PrinterProcessingStateDetail_InserterWarmingUp,
		"interlockclosed":                           PrinterProcessingStateDetail_InterlockClosed,
		"interlockopen":                             PrinterProcessingStateDetail_InterlockOpen,
		"interpretercartridgeadded":                 PrinterProcessingStateDetail_InterpreterCartridgeAdded,
		"interpretercartridgedeleted":               PrinterProcessingStateDetail_InterpreterCartridgeDeleted,
		"interpretercomplexpageencountered":         PrinterProcessingStateDetail_InterpreterComplexPageEncountered,
		"interpretermemorydecrease":                 PrinterProcessingStateDetail_InterpreterMemoryDecrease,
		"interpretermemoryincrease":                 PrinterProcessingStateDetail_InterpreterMemoryIncrease,
		"interpreterresourceadded":                  PrinterProcessingStateDetail_InterpreterResourceAdded,
		"interpreterresourcedeleted":                PrinterProcessingStateDetail_InterpreterResourceDeleted,
		"interpreterresourceunavailable":            PrinterProcessingStateDetail_InterpreterResourceUnavailable,
		"lampateol":                                 PrinterProcessingStateDetail_LampAtEol,
		"lampfailure":                               PrinterProcessingStateDetail_LampFailure,
		"lampneareol":                               PrinterProcessingStateDetail_LampNearEol,
		"laserateol":                                PrinterProcessingStateDetail_LaserAtEol,
		"laserfailure":                              PrinterProcessingStateDetail_LaserFailure,
		"laserneareol":                              PrinterProcessingStateDetail_LaserNearEol,
		"makeenvelopeadded":                         PrinterProcessingStateDetail_MakeEnvelopeAdded,
		"makeenvelopealmostempty":                   PrinterProcessingStateDetail_MakeEnvelopeAlmostEmpty,
		"makeenvelopealmostfull":                    PrinterProcessingStateDetail_MakeEnvelopeAlmostFull,
		"makeenvelopeatlimit":                       PrinterProcessingStateDetail_MakeEnvelopeAtLimit,
		"makeenvelopeclosed":                        PrinterProcessingStateDetail_MakeEnvelopeClosed,
		"makeenvelopeconfigurationchange":           PrinterProcessingStateDetail_MakeEnvelopeConfigurationChange,
		"makeenvelopecoverclosed":                   PrinterProcessingStateDetail_MakeEnvelopeCoverClosed,
		"makeenvelopecoveropen":                     PrinterProcessingStateDetail_MakeEnvelopeCoverOpen,
		"makeenvelopeempty":                         PrinterProcessingStateDetail_MakeEnvelopeEmpty,
		"makeenvelopefull":                          PrinterProcessingStateDetail_MakeEnvelopeFull,
		"makeenvelopeinterlockclosed":               PrinterProcessingStateDetail_MakeEnvelopeInterlockClosed,
		"makeenvelopeinterlockopen":                 PrinterProcessingStateDetail_MakeEnvelopeInterlockOpen,
		"makeenvelopejam":                           PrinterProcessingStateDetail_MakeEnvelopeJam,
		"makeenvelopelifealmostover":                PrinterProcessingStateDetail_MakeEnvelopeLifeAlmostOver,
		"makeenvelopelifeover":                      PrinterProcessingStateDetail_MakeEnvelopeLifeOver,
		"makeenvelopememoryexhausted":               PrinterProcessingStateDetail_MakeEnvelopeMemoryExhausted,
		"makeenvelopemissing":                       PrinterProcessingStateDetail_MakeEnvelopeMissing,
		"makeenvelopemotorfailure":                  PrinterProcessingStateDetail_MakeEnvelopeMotorFailure,
		"makeenvelopenearlimit":                     PrinterProcessingStateDetail_MakeEnvelopeNearLimit,
		"makeenvelopeoffline":                       PrinterProcessingStateDetail_MakeEnvelopeOffline,
		"makeenvelopeopened":                        PrinterProcessingStateDetail_MakeEnvelopeOpened,
		"makeenvelopeovertemperature":               PrinterProcessingStateDetail_MakeEnvelopeOverTemperature,
		"makeenvelopepowersaver":                    PrinterProcessingStateDetail_MakeEnvelopePowerSaver,
		"makeenveloperecoverablefailure":            PrinterProcessingStateDetail_MakeEnvelopeRecoverableFailure,
		"makeenveloperecoverablestorage":            PrinterProcessingStateDetail_MakeEnvelopeRecoverableStorage,
		"makeenveloperemoved":                       PrinterProcessingStateDetail_MakeEnvelopeRemoved,
		"makeenveloperesourceadded":                 PrinterProcessingStateDetail_MakeEnvelopeResourceAdded,
		"makeenveloperesourceremoved":               PrinterProcessingStateDetail_MakeEnvelopeResourceRemoved,
		"makeenvelopethermistorfailure":             PrinterProcessingStateDetail_MakeEnvelopeThermistorFailure,
		"makeenvelopetimingfailure":                 PrinterProcessingStateDetail_MakeEnvelopeTimingFailure,
		"makeenvelopeturnedoff":                     PrinterProcessingStateDetail_MakeEnvelopeTurnedOff,
		"makeenvelopeturnedon":                      PrinterProcessingStateDetail_MakeEnvelopeTurnedOn,
		"makeenvelopeundertemperature":              PrinterProcessingStateDetail_MakeEnvelopeUnderTemperature,
		"makeenvelopeunrecoverablefailure":          PrinterProcessingStateDetail_MakeEnvelopeUnrecoverableFailure,
		"makeenvelopeunrecoverablestorageerror":     PrinterProcessingStateDetail_MakeEnvelopeUnrecoverableStorageError,
		"makeenvelopewarmingup":                     PrinterProcessingStateDetail_MakeEnvelopeWarmingUp,
		"markeradjustingprintquality":               PrinterProcessingStateDetail_MarkerAdjustingPrintQuality,
		"markercleanermissing":                      PrinterProcessingStateDetail_MarkerCleanerMissing,
		"markerdeveloperalmostempty":                PrinterProcessingStateDetail_MarkerDeveloperAlmostEmpty,
		"markerdeveloperempty":                      PrinterProcessingStateDetail_MarkerDeveloperEmpty,
		"markerdevelopermissing":                    PrinterProcessingStateDetail_MarkerDeveloperMissing,
		"markerfusermissing":                        PrinterProcessingStateDetail_MarkerFuserMissing,
		"markerfuserthermistorfailure":              PrinterProcessingStateDetail_MarkerFuserThermistorFailure,
		"markerfusertimingfailure":                  PrinterProcessingStateDetail_MarkerFuserTimingFailure,
		"markerinkalmostempty":                      PrinterProcessingStateDetail_MarkerInkAlmostEmpty,
		"markerinkempty":                            PrinterProcessingStateDetail_MarkerInkEmpty,
		"markerinkmissing":                          PrinterProcessingStateDetail_MarkerInkMissing,
		"markeropcmissing":                          PrinterProcessingStateDetail_MarkerOpcMissing,
		"markerprintribbonalmostempty":              PrinterProcessingStateDetail_MarkerPrintRibbonAlmostEmpty,
		"markerprintribbonempty":                    PrinterProcessingStateDetail_MarkerPrintRibbonEmpty,
		"markerprintribbonmissing":                  PrinterProcessingStateDetail_MarkerPrintRibbonMissing,
		"markersupplyalmostempty":                   PrinterProcessingStateDetail_MarkerSupplyAlmostEmpty,
		"markersupplyempty":                         PrinterProcessingStateDetail_MarkerSupplyEmpty,
		"markersupplylow":                           PrinterProcessingStateDetail_MarkerSupplyLow,
		"markersupplymissing":                       PrinterProcessingStateDetail_MarkerSupplyMissing,
		"markertonercartridgemissing":               PrinterProcessingStateDetail_MarkerTonerCartridgeMissing,
		"markertonermissing":                        PrinterProcessingStateDetail_MarkerTonerMissing,
		"markerwastealmostfull":                     PrinterProcessingStateDetail_MarkerWasteAlmostFull,
		"markerwastefull":                           PrinterProcessingStateDetail_MarkerWasteFull,
		"markerwasteinkreceptaclealmostfull":        PrinterProcessingStateDetail_MarkerWasteInkReceptacleAlmostFull,
		"markerwasteinkreceptaclefull":              PrinterProcessingStateDetail_MarkerWasteInkReceptacleFull,
		"markerwasteinkreceptaclemissing":           PrinterProcessingStateDetail_MarkerWasteInkReceptacleMissing,
		"markerwastemissing":                        PrinterProcessingStateDetail_MarkerWasteMissing,
		"markerwastetonerreceptaclealmostfull":      PrinterProcessingStateDetail_MarkerWasteTonerReceptacleAlmostFull,
		"markerwastetonerreceptaclefull":            PrinterProcessingStateDetail_MarkerWasteTonerReceptacleFull,
		"markerwastetonerreceptaclemissing":         PrinterProcessingStateDetail_MarkerWasteTonerReceptacleMissing,
		"materialempty":                             PrinterProcessingStateDetail_MaterialEmpty,
		"materiallow":                               PrinterProcessingStateDetail_MaterialLow,
		"materialneeded":                            PrinterProcessingStateDetail_MaterialNeeded,
		"mediadrying":                               PrinterProcessingStateDetail_MediaDrying,
		"mediaempty":                                PrinterProcessingStateDetail_MediaEmpty,
		"mediajam":                                  PrinterProcessingStateDetail_MediaJam,
		"medialow":                                  PrinterProcessingStateDetail_MediaLow,
		"medianeeded":                               PrinterProcessingStateDetail_MediaNeeded,
		"mediapathcannotduplexmediaselected":        PrinterProcessingStateDetail_MediaPathCannotDuplexMediaSelected,
		"mediapathfailure":                          PrinterProcessingStateDetail_MediaPathFailure,
		"mediapathinputempty":                       PrinterProcessingStateDetail_MediaPathInputEmpty,
		"mediapathinputfeederror":                   PrinterProcessingStateDetail_MediaPathInputFeedError,
		"mediapathinputjam":                         PrinterProcessingStateDetail_MediaPathInputJam,
		"mediapathinputrequest":                     PrinterProcessingStateDetail_MediaPathInputRequest,
		"mediapathjam":                              PrinterProcessingStateDetail_MediaPathJam,
		"mediapathmediatrayalmostfull":              PrinterProcessingStateDetail_MediaPathMediaTrayAlmostFull,
		"mediapathmediatrayfull":                    PrinterProcessingStateDetail_MediaPathMediaTrayFull,
		"mediapathmediatraymissing":                 PrinterProcessingStateDetail_MediaPathMediaTrayMissing,
		"mediapathoutputfeederror":                  PrinterProcessingStateDetail_MediaPathOutputFeedError,
		"mediapathoutputfull":                       PrinterProcessingStateDetail_MediaPathOutputFull,
		"mediapathoutputjam":                        PrinterProcessingStateDetail_MediaPathOutputJam,
		"mediapathpickrollerfailure":                PrinterProcessingStateDetail_MediaPathPickRollerFailure,
		"mediapathpickrollerlifeover":               PrinterProcessingStateDetail_MediaPathPickRollerLifeOver,
		"mediapathpickrollerlifewarn":               PrinterProcessingStateDetail_MediaPathPickRollerLifeWarn,
		"mediapathpickrollermissing":                PrinterProcessingStateDetail_MediaPathPickRollerMissing,
		"motorfailure":                              PrinterProcessingStateDetail_MotorFailure,
		"movingtopaused":                            PrinterProcessingStateDetail_MovingToPaused,
		"none":                                      PrinterProcessingStateDetail_None,
		"opticalphotoconductorlifeover":             PrinterProcessingStateDetail_OpticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife":        PrinterProcessingStateDetail_OpticalPhotoConductorNearEndOfLife,
		"other":                                     PrinterProcessingStateDetail_Other,
		"outputareaalmostfull":                      PrinterProcessingStateDetail_OutputAreaAlmostFull,
		"outputareafull":                            PrinterProcessingStateDetail_OutputAreaFull,
		"outputmailboxselectfailure":                PrinterProcessingStateDetail_OutputMailboxSelectFailure,
		"outputmediatrayfailure":                    PrinterProcessingStateDetail_OutputMediaTrayFailure,
		"outputmediatrayfeederror":                  PrinterProcessingStateDetail_OutputMediaTrayFeedError,
		"outputmediatrayjam":                        PrinterProcessingStateDetail_OutputMediaTrayJam,
		"outputtraymissing":                         PrinterProcessingStateDetail_OutputTrayMissing,
		"paused":                                    PrinterProcessingStateDetail_Paused,
		"perforateradded":                           PrinterProcessingStateDetail_PerforaterAdded,
		"perforateralmostempty":                     PrinterProcessingStateDetail_PerforaterAlmostEmpty,
		"perforateralmostfull":                      PrinterProcessingStateDetail_PerforaterAlmostFull,
		"perforateratlimit":                         PrinterProcessingStateDetail_PerforaterAtLimit,
		"perforaterclosed":                          PrinterProcessingStateDetail_PerforaterClosed,
		"perforaterconfigurationchange":             PrinterProcessingStateDetail_PerforaterConfigurationChange,
		"perforatercoverclosed":                     PrinterProcessingStateDetail_PerforaterCoverClosed,
		"perforatercoveropen":                       PrinterProcessingStateDetail_PerforaterCoverOpen,
		"perforaterempty":                           PrinterProcessingStateDetail_PerforaterEmpty,
		"perforaterfull":                            PrinterProcessingStateDetail_PerforaterFull,
		"perforaterinterlockclosed":                 PrinterProcessingStateDetail_PerforaterInterlockClosed,
		"perforaterinterlockopen":                   PrinterProcessingStateDetail_PerforaterInterlockOpen,
		"perforaterjam":                             PrinterProcessingStateDetail_PerforaterJam,
		"perforaterlifealmostover":                  PrinterProcessingStateDetail_PerforaterLifeAlmostOver,
		"perforaterlifeover":                        PrinterProcessingStateDetail_PerforaterLifeOver,
		"perforatermemoryexhausted":                 PrinterProcessingStateDetail_PerforaterMemoryExhausted,
		"perforatermissing":                         PrinterProcessingStateDetail_PerforaterMissing,
		"perforatermotorfailure":                    PrinterProcessingStateDetail_PerforaterMotorFailure,
		"perforaternearlimit":                       PrinterProcessingStateDetail_PerforaterNearLimit,
		"perforateroffline":                         PrinterProcessingStateDetail_PerforaterOffline,
		"perforateropened":                          PrinterProcessingStateDetail_PerforaterOpened,
		"perforaterovertemperature":                 PrinterProcessingStateDetail_PerforaterOverTemperature,
		"perforaterpowersaver":                      PrinterProcessingStateDetail_PerforaterPowerSaver,
		"perforaterrecoverablefailure":              PrinterProcessingStateDetail_PerforaterRecoverableFailure,
		"perforaterrecoverablestorage":              PrinterProcessingStateDetail_PerforaterRecoverableStorage,
		"perforaterremoved":                         PrinterProcessingStateDetail_PerforaterRemoved,
		"perforaterresourceadded":                   PrinterProcessingStateDetail_PerforaterResourceAdded,
		"perforaterresourceremoved":                 PrinterProcessingStateDetail_PerforaterResourceRemoved,
		"perforaterthermistorfailure":               PrinterProcessingStateDetail_PerforaterThermistorFailure,
		"perforatertimingfailure":                   PrinterProcessingStateDetail_PerforaterTimingFailure,
		"perforaterturnedoff":                       PrinterProcessingStateDetail_PerforaterTurnedOff,
		"perforaterturnedon":                        PrinterProcessingStateDetail_PerforaterTurnedOn,
		"perforaterundertemperature":                PrinterProcessingStateDetail_PerforaterUnderTemperature,
		"perforaterunrecoverablefailure":            PrinterProcessingStateDetail_PerforaterUnrecoverableFailure,
		"perforaterunrecoverablestorageerror":       PrinterProcessingStateDetail_PerforaterUnrecoverableStorageError,
		"perforaterwarmingup":                       PrinterProcessingStateDetail_PerforaterWarmingUp,
		"platformcooling":                           PrinterProcessingStateDetail_PlatformCooling,
		"platformfailure":                           PrinterProcessingStateDetail_PlatformFailure,
		"platformheating":                           PrinterProcessingStateDetail_PlatformHeating,
		"platformtemperaturehigh":                   PrinterProcessingStateDetail_PlatformTemperatureHigh,
		"platformtemperaturelow":                    PrinterProcessingStateDetail_PlatformTemperatureLow,
		"powerdown":                                 PrinterProcessingStateDetail_PowerDown,
		"powerup":                                   PrinterProcessingStateDetail_PowerUp,
		"printermanualreset":                        PrinterProcessingStateDetail_PrinterManualReset,
		"printernmsreset":                           PrinterProcessingStateDetail_PrinterNmsReset,
		"printerreadytoprint":                       PrinterProcessingStateDetail_PrinterReadyToPrint,
		"puncheradded":                              PrinterProcessingStateDetail_PuncherAdded,
		"puncheralmostempty":                        PrinterProcessingStateDetail_PuncherAlmostEmpty,
		"puncheralmostfull":                         PrinterProcessingStateDetail_PuncherAlmostFull,
		"puncheratlimit":                            PrinterProcessingStateDetail_PuncherAtLimit,
		"puncherclosed":                             PrinterProcessingStateDetail_PuncherClosed,
		"puncherconfigurationchange":                PrinterProcessingStateDetail_PuncherConfigurationChange,
		"punchercoverclosed":                        PrinterProcessingStateDetail_PuncherCoverClosed,
		"punchercoveropen":                          PrinterProcessingStateDetail_PuncherCoverOpen,
		"puncherempty":                              PrinterProcessingStateDetail_PuncherEmpty,
		"puncherfull":                               PrinterProcessingStateDetail_PuncherFull,
		"puncherinterlockclosed":                    PrinterProcessingStateDetail_PuncherInterlockClosed,
		"puncherinterlockopen":                      PrinterProcessingStateDetail_PuncherInterlockOpen,
		"puncherjam":                                PrinterProcessingStateDetail_PuncherJam,
		"puncherlifealmostover":                     PrinterProcessingStateDetail_PuncherLifeAlmostOver,
		"puncherlifeover":                           PrinterProcessingStateDetail_PuncherLifeOver,
		"punchermemoryexhausted":                    PrinterProcessingStateDetail_PuncherMemoryExhausted,
		"punchermissing":                            PrinterProcessingStateDetail_PuncherMissing,
		"punchermotorfailure":                       PrinterProcessingStateDetail_PuncherMotorFailure,
		"punchernearlimit":                          PrinterProcessingStateDetail_PuncherNearLimit,
		"puncheroffline":                            PrinterProcessingStateDetail_PuncherOffline,
		"puncheropened":                             PrinterProcessingStateDetail_PuncherOpened,
		"puncherovertemperature":                    PrinterProcessingStateDetail_PuncherOverTemperature,
		"puncherpowersaver":                         PrinterProcessingStateDetail_PuncherPowerSaver,
		"puncherrecoverablefailure":                 PrinterProcessingStateDetail_PuncherRecoverableFailure,
		"puncherrecoverablestorage":                 PrinterProcessingStateDetail_PuncherRecoverableStorage,
		"puncherremoved":                            PrinterProcessingStateDetail_PuncherRemoved,
		"puncherresourceadded":                      PrinterProcessingStateDetail_PuncherResourceAdded,
		"puncherresourceremoved":                    PrinterProcessingStateDetail_PuncherResourceRemoved,
		"puncherthermistorfailure":                  PrinterProcessingStateDetail_PuncherThermistorFailure,
		"punchertimingfailure":                      PrinterProcessingStateDetail_PuncherTimingFailure,
		"puncherturnedoff":                          PrinterProcessingStateDetail_PuncherTurnedOff,
		"puncherturnedon":                           PrinterProcessingStateDetail_PuncherTurnedOn,
		"puncherundertemperature":                   PrinterProcessingStateDetail_PuncherUnderTemperature,
		"puncherunrecoverablefailure":               PrinterProcessingStateDetail_PuncherUnrecoverableFailure,
		"puncherunrecoverablestorageerror":          PrinterProcessingStateDetail_PuncherUnrecoverableStorageError,
		"puncherwarmingup":                          PrinterProcessingStateDetail_PuncherWarmingUp,
		"resuming":                                  PrinterProcessingStateDetail_Resuming,
		"scanmediapathfailure":                      PrinterProcessingStateDetail_ScanMediaPathFailure,
		"scanmediapathinputempty":                   PrinterProcessingStateDetail_ScanMediaPathInputEmpty,
		"scanmediapathinputfeederror":               PrinterProcessingStateDetail_ScanMediaPathInputFeedError,
		"scanmediapathinputjam":                     PrinterProcessingStateDetail_ScanMediaPathInputJam,
		"scanmediapathinputrequest":                 PrinterProcessingStateDetail_ScanMediaPathInputRequest,
		"scanmediapathjam":                          PrinterProcessingStateDetail_ScanMediaPathJam,
		"scanmediapathoutputfeederror":              PrinterProcessingStateDetail_ScanMediaPathOutputFeedError,
		"scanmediapathoutputfull":                   PrinterProcessingStateDetail_ScanMediaPathOutputFull,
		"scanmediapathoutputjam":                    PrinterProcessingStateDetail_ScanMediaPathOutputJam,
		"scanmediapathpickrollerfailure":            PrinterProcessingStateDetail_ScanMediaPathPickRollerFailure,
		"scanmediapathpickrollerlifeover":           PrinterProcessingStateDetail_ScanMediaPathPickRollerLifeOver,
		"scanmediapathpickrollerlifewarn":           PrinterProcessingStateDetail_ScanMediaPathPickRollerLifeWarn,
		"scanmediapathpickrollermissing":            PrinterProcessingStateDetail_ScanMediaPathPickRollerMissing,
		"scanmediapathtrayalmostfull":               PrinterProcessingStateDetail_ScanMediaPathTrayAlmostFull,
		"scanmediapathtrayfull":                     PrinterProcessingStateDetail_ScanMediaPathTrayFull,
		"scanmediapathtraymissing":                  PrinterProcessingStateDetail_ScanMediaPathTrayMissing,
		"scannerlightfailure":                       PrinterProcessingStateDetail_ScannerLightFailure,
		"scannerlightlifealmostover":                PrinterProcessingStateDetail_ScannerLightLifeAlmostOver,
		"scannerlightlifeover":                      PrinterProcessingStateDetail_ScannerLightLifeOver,
		"scannerlightmissing":                       PrinterProcessingStateDetail_ScannerLightMissing,
		"scannersensorfailure":                      PrinterProcessingStateDetail_ScannerSensorFailure,
		"scannersensorlifealmostover":               PrinterProcessingStateDetail_ScannerSensorLifeAlmostOver,
		"scannersensorlifeover":                     PrinterProcessingStateDetail_ScannerSensorLifeOver,
		"scannersensormissing":                      PrinterProcessingStateDetail_ScannerSensorMissing,
		"separationcutteradded":                     PrinterProcessingStateDetail_SeparationCutterAdded,
		"separationcutteralmostempty":               PrinterProcessingStateDetail_SeparationCutterAlmostEmpty,
		"separationcutteralmostfull":                PrinterProcessingStateDetail_SeparationCutterAlmostFull,
		"separationcutteratlimit":                   PrinterProcessingStateDetail_SeparationCutterAtLimit,
		"separationcutterclosed":                    PrinterProcessingStateDetail_SeparationCutterClosed,
		"separationcutterconfigurationchange":       PrinterProcessingStateDetail_SeparationCutterConfigurationChange,
		"separationcuttercoverclosed":               PrinterProcessingStateDetail_SeparationCutterCoverClosed,
		"separationcuttercoveropen":                 PrinterProcessingStateDetail_SeparationCutterCoverOpen,
		"separationcutterempty":                     PrinterProcessingStateDetail_SeparationCutterEmpty,
		"separationcutterfull":                      PrinterProcessingStateDetail_SeparationCutterFull,
		"separationcutterinterlockclosed":           PrinterProcessingStateDetail_SeparationCutterInterlockClosed,
		"separationcutterinterlockopen":             PrinterProcessingStateDetail_SeparationCutterInterlockOpen,
		"separationcutterjam":                       PrinterProcessingStateDetail_SeparationCutterJam,
		"separationcutterlifealmostover":            PrinterProcessingStateDetail_SeparationCutterLifeAlmostOver,
		"separationcutterlifeover":                  PrinterProcessingStateDetail_SeparationCutterLifeOver,
		"separationcuttermemoryexhausted":           PrinterProcessingStateDetail_SeparationCutterMemoryExhausted,
		"separationcuttermissing":                   PrinterProcessingStateDetail_SeparationCutterMissing,
		"separationcuttermotorfailure":              PrinterProcessingStateDetail_SeparationCutterMotorFailure,
		"separationcutternearlimit":                 PrinterProcessingStateDetail_SeparationCutterNearLimit,
		"separationcutteroffline":                   PrinterProcessingStateDetail_SeparationCutterOffline,
		"separationcutteropened":                    PrinterProcessingStateDetail_SeparationCutterOpened,
		"separationcutterovertemperature":           PrinterProcessingStateDetail_SeparationCutterOverTemperature,
		"separationcutterpowersaver":                PrinterProcessingStateDetail_SeparationCutterPowerSaver,
		"separationcutterrecoverablefailure":        PrinterProcessingStateDetail_SeparationCutterRecoverableFailure,
		"separationcutterrecoverablestorage":        PrinterProcessingStateDetail_SeparationCutterRecoverableStorage,
		"separationcutterremoved":                   PrinterProcessingStateDetail_SeparationCutterRemoved,
		"separationcutterresourceadded":             PrinterProcessingStateDetail_SeparationCutterResourceAdded,
		"separationcutterresourceremoved":           PrinterProcessingStateDetail_SeparationCutterResourceRemoved,
		"separationcutterthermistorfailure":         PrinterProcessingStateDetail_SeparationCutterThermistorFailure,
		"separationcuttertimingfailure":             PrinterProcessingStateDetail_SeparationCutterTimingFailure,
		"separationcutterturnedoff":                 PrinterProcessingStateDetail_SeparationCutterTurnedOff,
		"separationcutterturnedon":                  PrinterProcessingStateDetail_SeparationCutterTurnedOn,
		"separationcutterundertemperature":          PrinterProcessingStateDetail_SeparationCutterUnderTemperature,
		"separationcutterunrecoverablefailure":      PrinterProcessingStateDetail_SeparationCutterUnrecoverableFailure,
		"separationcutterunrecoverablestorageerror": PrinterProcessingStateDetail_SeparationCutterUnrecoverableStorageError,
		"separationcutterwarmingup":                 PrinterProcessingStateDetail_SeparationCutterWarmingUp,
		"sheetrotatoradded":                         PrinterProcessingStateDetail_SheetRotatorAdded,
		"sheetrotatoralmostempty":                   PrinterProcessingStateDetail_SheetRotatorAlmostEmpty,
		"sheetrotatoralmostfull":                    PrinterProcessingStateDetail_SheetRotatorAlmostFull,
		"sheetrotatoratlimit":                       PrinterProcessingStateDetail_SheetRotatorAtLimit,
		"sheetrotatorclosed":                        PrinterProcessingStateDetail_SheetRotatorClosed,
		"sheetrotatorconfigurationchange":           PrinterProcessingStateDetail_SheetRotatorConfigurationChange,
		"sheetrotatorcoverclosed":                   PrinterProcessingStateDetail_SheetRotatorCoverClosed,
		"sheetrotatorcoveropen":                     PrinterProcessingStateDetail_SheetRotatorCoverOpen,
		"sheetrotatorempty":                         PrinterProcessingStateDetail_SheetRotatorEmpty,
		"sheetrotatorfull":                          PrinterProcessingStateDetail_SheetRotatorFull,
		"sheetrotatorinterlockclosed":               PrinterProcessingStateDetail_SheetRotatorInterlockClosed,
		"sheetrotatorinterlockopen":                 PrinterProcessingStateDetail_SheetRotatorInterlockOpen,
		"sheetrotatorjam":                           PrinterProcessingStateDetail_SheetRotatorJam,
		"sheetrotatorlifealmostover":                PrinterProcessingStateDetail_SheetRotatorLifeAlmostOver,
		"sheetrotatorlifeover":                      PrinterProcessingStateDetail_SheetRotatorLifeOver,
		"sheetrotatormemoryexhausted":               PrinterProcessingStateDetail_SheetRotatorMemoryExhausted,
		"sheetrotatormissing":                       PrinterProcessingStateDetail_SheetRotatorMissing,
		"sheetrotatormotorfailure":                  PrinterProcessingStateDetail_SheetRotatorMotorFailure,
		"sheetrotatornearlimit":                     PrinterProcessingStateDetail_SheetRotatorNearLimit,
		"sheetrotatoroffline":                       PrinterProcessingStateDetail_SheetRotatorOffline,
		"sheetrotatoropened":                        PrinterProcessingStateDetail_SheetRotatorOpened,
		"sheetrotatorovertemperature":               PrinterProcessingStateDetail_SheetRotatorOverTemperature,
		"sheetrotatorpowersaver":                    PrinterProcessingStateDetail_SheetRotatorPowerSaver,
		"sheetrotatorrecoverablefailure":            PrinterProcessingStateDetail_SheetRotatorRecoverableFailure,
		"sheetrotatorrecoverablestorage":            PrinterProcessingStateDetail_SheetRotatorRecoverableStorage,
		"sheetrotatorremoved":                       PrinterProcessingStateDetail_SheetRotatorRemoved,
		"sheetrotatorresourceadded":                 PrinterProcessingStateDetail_SheetRotatorResourceAdded,
		"sheetrotatorresourceremoved":               PrinterProcessingStateDetail_SheetRotatorResourceRemoved,
		"sheetrotatorthermistorfailure":             PrinterProcessingStateDetail_SheetRotatorThermistorFailure,
		"sheetrotatortimingfailure":                 PrinterProcessingStateDetail_SheetRotatorTimingFailure,
		"sheetrotatorturnedoff":                     PrinterProcessingStateDetail_SheetRotatorTurnedOff,
		"sheetrotatorturnedon":                      PrinterProcessingStateDetail_SheetRotatorTurnedOn,
		"sheetrotatorundertemperature":              PrinterProcessingStateDetail_SheetRotatorUnderTemperature,
		"sheetrotatorunrecoverablefailure":          PrinterProcessingStateDetail_SheetRotatorUnrecoverableFailure,
		"sheetrotatorunrecoverablestorageerror":     PrinterProcessingStateDetail_SheetRotatorUnrecoverableStorageError,
		"sheetrotatorwarmingup":                     PrinterProcessingStateDetail_SheetRotatorWarmingUp,
		"shutdown":                                  PrinterProcessingStateDetail_Shutdown,
		"slitteradded":                              PrinterProcessingStateDetail_SlitterAdded,
		"slitteralmostempty":                        PrinterProcessingStateDetail_SlitterAlmostEmpty,
		"slitteralmostfull":                         PrinterProcessingStateDetail_SlitterAlmostFull,
		"slitteratlimit":                            PrinterProcessingStateDetail_SlitterAtLimit,
		"slitterclosed":                             PrinterProcessingStateDetail_SlitterClosed,
		"slitterconfigurationchange":                PrinterProcessingStateDetail_SlitterConfigurationChange,
		"slittercoverclosed":                        PrinterProcessingStateDetail_SlitterCoverClosed,
		"slittercoveropen":                          PrinterProcessingStateDetail_SlitterCoverOpen,
		"slitterempty":                              PrinterProcessingStateDetail_SlitterEmpty,
		"slitterfull":                               PrinterProcessingStateDetail_SlitterFull,
		"slitterinterlockclosed":                    PrinterProcessingStateDetail_SlitterInterlockClosed,
		"slitterinterlockopen":                      PrinterProcessingStateDetail_SlitterInterlockOpen,
		"slitterjam":                                PrinterProcessingStateDetail_SlitterJam,
		"slitterlifealmostover":                     PrinterProcessingStateDetail_SlitterLifeAlmostOver,
		"slitterlifeover":                           PrinterProcessingStateDetail_SlitterLifeOver,
		"slittermemoryexhausted":                    PrinterProcessingStateDetail_SlitterMemoryExhausted,
		"slittermissing":                            PrinterProcessingStateDetail_SlitterMissing,
		"slittermotorfailure":                       PrinterProcessingStateDetail_SlitterMotorFailure,
		"slitternearlimit":                          PrinterProcessingStateDetail_SlitterNearLimit,
		"slitteroffline":                            PrinterProcessingStateDetail_SlitterOffline,
		"slitteropened":                             PrinterProcessingStateDetail_SlitterOpened,
		"slitterovertemperature":                    PrinterProcessingStateDetail_SlitterOverTemperature,
		"slitterpowersaver":                         PrinterProcessingStateDetail_SlitterPowerSaver,
		"slitterrecoverablefailure":                 PrinterProcessingStateDetail_SlitterRecoverableFailure,
		"slitterrecoverablestorage":                 PrinterProcessingStateDetail_SlitterRecoverableStorage,
		"slitterremoved":                            PrinterProcessingStateDetail_SlitterRemoved,
		"slitterresourceadded":                      PrinterProcessingStateDetail_SlitterResourceAdded,
		"slitterresourceremoved":                    PrinterProcessingStateDetail_SlitterResourceRemoved,
		"slitterthermistorfailure":                  PrinterProcessingStateDetail_SlitterThermistorFailure,
		"slittertimingfailure":                      PrinterProcessingStateDetail_SlitterTimingFailure,
		"slitterturnedoff":                          PrinterProcessingStateDetail_SlitterTurnedOff,
		"slitterturnedon":                           PrinterProcessingStateDetail_SlitterTurnedOn,
		"slitterundertemperature":                   PrinterProcessingStateDetail_SlitterUnderTemperature,
		"slitterunrecoverablefailure":               PrinterProcessingStateDetail_SlitterUnrecoverableFailure,
		"slitterunrecoverablestorageerror":          PrinterProcessingStateDetail_SlitterUnrecoverableStorageError,
		"slitterwarmingup":                          PrinterProcessingStateDetail_SlitterWarmingUp,
		"spoolareafull":                             PrinterProcessingStateDetail_SpoolAreaFull,
		"stackeradded":                              PrinterProcessingStateDetail_StackerAdded,
		"stackeralmostempty":                        PrinterProcessingStateDetail_StackerAlmostEmpty,
		"stackeralmostfull":                         PrinterProcessingStateDetail_StackerAlmostFull,
		"stackeratlimit":                            PrinterProcessingStateDetail_StackerAtLimit,
		"stackerclosed":                             PrinterProcessingStateDetail_StackerClosed,
		"stackerconfigurationchange":                PrinterProcessingStateDetail_StackerConfigurationChange,
		"stackercoverclosed":                        PrinterProcessingStateDetail_StackerCoverClosed,
		"stackercoveropen":                          PrinterProcessingStateDetail_StackerCoverOpen,
		"stackerempty":                              PrinterProcessingStateDetail_StackerEmpty,
		"stackerfull":                               PrinterProcessingStateDetail_StackerFull,
		"stackerinterlockclosed":                    PrinterProcessingStateDetail_StackerInterlockClosed,
		"stackerinterlockopen":                      PrinterProcessingStateDetail_StackerInterlockOpen,
		"stackerjam":                                PrinterProcessingStateDetail_StackerJam,
		"stackerlifealmostover":                     PrinterProcessingStateDetail_StackerLifeAlmostOver,
		"stackerlifeover":                           PrinterProcessingStateDetail_StackerLifeOver,
		"stackermemoryexhausted":                    PrinterProcessingStateDetail_StackerMemoryExhausted,
		"stackermissing":                            PrinterProcessingStateDetail_StackerMissing,
		"stackermotorfailure":                       PrinterProcessingStateDetail_StackerMotorFailure,
		"stackernearlimit":                          PrinterProcessingStateDetail_StackerNearLimit,
		"stackeroffline":                            PrinterProcessingStateDetail_StackerOffline,
		"stackeropened":                             PrinterProcessingStateDetail_StackerOpened,
		"stackerovertemperature":                    PrinterProcessingStateDetail_StackerOverTemperature,
		"stackerpowersaver":                         PrinterProcessingStateDetail_StackerPowerSaver,
		"stackerrecoverablefailure":                 PrinterProcessingStateDetail_StackerRecoverableFailure,
		"stackerrecoverablestorage":                 PrinterProcessingStateDetail_StackerRecoverableStorage,
		"stackerremoved":                            PrinterProcessingStateDetail_StackerRemoved,
		"stackerresourceadded":                      PrinterProcessingStateDetail_StackerResourceAdded,
		"stackerresourceremoved":                    PrinterProcessingStateDetail_StackerResourceRemoved,
		"stackerthermistorfailure":                  PrinterProcessingStateDetail_StackerThermistorFailure,
		"stackertimingfailure":                      PrinterProcessingStateDetail_StackerTimingFailure,
		"stackerturnedoff":                          PrinterProcessingStateDetail_StackerTurnedOff,
		"stackerturnedon":                           PrinterProcessingStateDetail_StackerTurnedOn,
		"stackerundertemperature":                   PrinterProcessingStateDetail_StackerUnderTemperature,
		"stackerunrecoverablefailure":               PrinterProcessingStateDetail_StackerUnrecoverableFailure,
		"stackerunrecoverablestorageerror":          PrinterProcessingStateDetail_StackerUnrecoverableStorageError,
		"stackerwarmingup":                          PrinterProcessingStateDetail_StackerWarmingUp,
		"standby":                                   PrinterProcessingStateDetail_Standby,
		"stapleradded":                              PrinterProcessingStateDetail_StaplerAdded,
		"stapleralmostempty":                        PrinterProcessingStateDetail_StaplerAlmostEmpty,
		"stapleralmostfull":                         PrinterProcessingStateDetail_StaplerAlmostFull,
		"stapleratlimit":                            PrinterProcessingStateDetail_StaplerAtLimit,
		"staplerclosed":                             PrinterProcessingStateDetail_StaplerClosed,
		"staplerconfigurationchange":                PrinterProcessingStateDetail_StaplerConfigurationChange,
		"staplercoverclosed":                        PrinterProcessingStateDetail_StaplerCoverClosed,
		"staplercoveropen":                          PrinterProcessingStateDetail_StaplerCoverOpen,
		"staplerempty":                              PrinterProcessingStateDetail_StaplerEmpty,
		"staplerfull":                               PrinterProcessingStateDetail_StaplerFull,
		"staplerinterlockclosed":                    PrinterProcessingStateDetail_StaplerInterlockClosed,
		"staplerinterlockopen":                      PrinterProcessingStateDetail_StaplerInterlockOpen,
		"staplerjam":                                PrinterProcessingStateDetail_StaplerJam,
		"staplerlifealmostover":                     PrinterProcessingStateDetail_StaplerLifeAlmostOver,
		"staplerlifeover":                           PrinterProcessingStateDetail_StaplerLifeOver,
		"staplermemoryexhausted":                    PrinterProcessingStateDetail_StaplerMemoryExhausted,
		"staplermissing":                            PrinterProcessingStateDetail_StaplerMissing,
		"staplermotorfailure":                       PrinterProcessingStateDetail_StaplerMotorFailure,
		"staplernearlimit":                          PrinterProcessingStateDetail_StaplerNearLimit,
		"stapleroffline":                            PrinterProcessingStateDetail_StaplerOffline,
		"stapleropened":                             PrinterProcessingStateDetail_StaplerOpened,
		"staplerovertemperature":                    PrinterProcessingStateDetail_StaplerOverTemperature,
		"staplerpowersaver":                         PrinterProcessingStateDetail_StaplerPowerSaver,
		"staplerrecoverablefailure":                 PrinterProcessingStateDetail_StaplerRecoverableFailure,
		"staplerrecoverablestorage":                 PrinterProcessingStateDetail_StaplerRecoverableStorage,
		"staplerremoved":                            PrinterProcessingStateDetail_StaplerRemoved,
		"staplerresourceadded":                      PrinterProcessingStateDetail_StaplerResourceAdded,
		"staplerresourceremoved":                    PrinterProcessingStateDetail_StaplerResourceRemoved,
		"staplerthermistorfailure":                  PrinterProcessingStateDetail_StaplerThermistorFailure,
		"staplertimingfailure":                      PrinterProcessingStateDetail_StaplerTimingFailure,
		"staplerturnedoff":                          PrinterProcessingStateDetail_StaplerTurnedOff,
		"staplerturnedon":                           PrinterProcessingStateDetail_StaplerTurnedOn,
		"staplerundertemperature":                   PrinterProcessingStateDetail_StaplerUnderTemperature,
		"staplerunrecoverablefailure":               PrinterProcessingStateDetail_StaplerUnrecoverableFailure,
		"staplerunrecoverablestorageerror":          PrinterProcessingStateDetail_StaplerUnrecoverableStorageError,
		"staplerwarmingup":                          PrinterProcessingStateDetail_StaplerWarmingUp,
		"stitcheradded":                             PrinterProcessingStateDetail_StitcherAdded,
		"stitcheralmostempty":                       PrinterProcessingStateDetail_StitcherAlmostEmpty,
		"stitcheralmostfull":                        PrinterProcessingStateDetail_StitcherAlmostFull,
		"stitcheratlimit":                           PrinterProcessingStateDetail_StitcherAtLimit,
		"stitcherclosed":                            PrinterProcessingStateDetail_StitcherClosed,
		"stitcherconfigurationchange":               PrinterProcessingStateDetail_StitcherConfigurationChange,
		"stitchercoverclosed":                       PrinterProcessingStateDetail_StitcherCoverClosed,
		"stitchercoveropen":                         PrinterProcessingStateDetail_StitcherCoverOpen,
		"stitcherempty":                             PrinterProcessingStateDetail_StitcherEmpty,
		"stitcherfull":                              PrinterProcessingStateDetail_StitcherFull,
		"stitcherinterlockclosed":                   PrinterProcessingStateDetail_StitcherInterlockClosed,
		"stitcherinterlockopen":                     PrinterProcessingStateDetail_StitcherInterlockOpen,
		"stitcherjam":                               PrinterProcessingStateDetail_StitcherJam,
		"stitcherlifealmostover":                    PrinterProcessingStateDetail_StitcherLifeAlmostOver,
		"stitcherlifeover":                          PrinterProcessingStateDetail_StitcherLifeOver,
		"stitchermemoryexhausted":                   PrinterProcessingStateDetail_StitcherMemoryExhausted,
		"stitchermissing":                           PrinterProcessingStateDetail_StitcherMissing,
		"stitchermotorfailure":                      PrinterProcessingStateDetail_StitcherMotorFailure,
		"stitchernearlimit":                         PrinterProcessingStateDetail_StitcherNearLimit,
		"stitcheroffline":                           PrinterProcessingStateDetail_StitcherOffline,
		"stitcheropened":                            PrinterProcessingStateDetail_StitcherOpened,
		"stitcherovertemperature":                   PrinterProcessingStateDetail_StitcherOverTemperature,
		"stitcherpowersaver":                        PrinterProcessingStateDetail_StitcherPowerSaver,
		"stitcherrecoverablefailure":                PrinterProcessingStateDetail_StitcherRecoverableFailure,
		"stitcherrecoverablestorage":                PrinterProcessingStateDetail_StitcherRecoverableStorage,
		"stitcherremoved":                           PrinterProcessingStateDetail_StitcherRemoved,
		"stitcherresourceadded":                     PrinterProcessingStateDetail_StitcherResourceAdded,
		"stitcherresourceremoved":                   PrinterProcessingStateDetail_StitcherResourceRemoved,
		"stitcherthermistorfailure":                 PrinterProcessingStateDetail_StitcherThermistorFailure,
		"stitchertimingfailure":                     PrinterProcessingStateDetail_StitcherTimingFailure,
		"stitcherturnedoff":                         PrinterProcessingStateDetail_StitcherTurnedOff,
		"stitcherturnedon":                          PrinterProcessingStateDetail_StitcherTurnedOn,
		"stitcherundertemperature":                  PrinterProcessingStateDetail_StitcherUnderTemperature,
		"stitcherunrecoverablefailure":              PrinterProcessingStateDetail_StitcherUnrecoverableFailure,
		"stitcherunrecoverablestorageerror":         PrinterProcessingStateDetail_StitcherUnrecoverableStorageError,
		"stitcherwarmingup":                         PrinterProcessingStateDetail_StitcherWarmingUp,
		"stoppedpartially":                          PrinterProcessingStateDetail_StoppedPartially,
		"stopping":                                  PrinterProcessingStateDetail_Stopping,
		"subunitadded":                              PrinterProcessingStateDetail_SubunitAdded,
		"subunitalmostempty":                        PrinterProcessingStateDetail_SubunitAlmostEmpty,
		"subunitalmostfull":                         PrinterProcessingStateDetail_SubunitAlmostFull,
		"subunitatlimit":                            PrinterProcessingStateDetail_SubunitAtLimit,
		"subunitclosed":                             PrinterProcessingStateDetail_SubunitClosed,
		"subunitcoolingdown":                        PrinterProcessingStateDetail_SubunitCoolingDown,
		"subunitempty":                              PrinterProcessingStateDetail_SubunitEmpty,
		"subunitfull":                               PrinterProcessingStateDetail_SubunitFull,
		"subunitlifealmostover":                     PrinterProcessingStateDetail_SubunitLifeAlmostOver,
		"subunitlifeover":                           PrinterProcessingStateDetail_SubunitLifeOver,
		"subunitmemoryexhausted":                    PrinterProcessingStateDetail_SubunitMemoryExhausted,
		"subunitmissing":                            PrinterProcessingStateDetail_SubunitMissing,
		"subunitmotorfailure":                       PrinterProcessingStateDetail_SubunitMotorFailure,
		"subunitnearlimit":                          PrinterProcessingStateDetail_SubunitNearLimit,
		"subunitoffline":                            PrinterProcessingStateDetail_SubunitOffline,
		"subunitopened":                             PrinterProcessingStateDetail_SubunitOpened,
		"subunitovertemperature":                    PrinterProcessingStateDetail_SubunitOverTemperature,
		"subunitpowersaver":                         PrinterProcessingStateDetail_SubunitPowerSaver,
		"subunitrecoverablefailure":                 PrinterProcessingStateDetail_SubunitRecoverableFailure,
		"subunitrecoverablestorage":                 PrinterProcessingStateDetail_SubunitRecoverableStorage,
		"subunitremoved":                            PrinterProcessingStateDetail_SubunitRemoved,
		"subunitresourceadded":                      PrinterProcessingStateDetail_SubunitResourceAdded,
		"subunitresourceremoved":                    PrinterProcessingStateDetail_SubunitResourceRemoved,
		"subunitthermistorfailure":                  PrinterProcessingStateDetail_SubunitThermistorFailure,
		"subunittimingfailure":                      PrinterProcessingStateDetail_SubunitTimingFailure,
		"subunitturnedoff":                          PrinterProcessingStateDetail_SubunitTurnedOff,
		"subunitturnedon":                           PrinterProcessingStateDetail_SubunitTurnedOn,
		"subunitundertemperature":                   PrinterProcessingStateDetail_SubunitUnderTemperature,
		"subunitunrecoverablefailure":               PrinterProcessingStateDetail_SubunitUnrecoverableFailure,
		"subunitunrecoverablestorage":               PrinterProcessingStateDetail_SubunitUnrecoverableStorage,
		"subunitwarmingup":                          PrinterProcessingStateDetail_SubunitWarmingUp,
		"suspend":                                   PrinterProcessingStateDetail_Suspend,
		"testing":                                   PrinterProcessingStateDetail_Testing,
		"timedout":                                  PrinterProcessingStateDetail_TimedOut,
		"tonerempty":                                PrinterProcessingStateDetail_TonerEmpty,
		"tonerlow":                                  PrinterProcessingStateDetail_TonerLow,
		"trimmeradded":                              PrinterProcessingStateDetail_TrimmerAdded,
		"trimmeralmostempty":                        PrinterProcessingStateDetail_TrimmerAlmostEmpty,
		"trimmeralmostfull":                         PrinterProcessingStateDetail_TrimmerAlmostFull,
		"trimmeratlimit":                            PrinterProcessingStateDetail_TrimmerAtLimit,
		"trimmerclosed":                             PrinterProcessingStateDetail_TrimmerClosed,
		"trimmerconfigurationchange":                PrinterProcessingStateDetail_TrimmerConfigurationChange,
		"trimmercoverclosed":                        PrinterProcessingStateDetail_TrimmerCoverClosed,
		"trimmercoveropen":                          PrinterProcessingStateDetail_TrimmerCoverOpen,
		"trimmerempty":                              PrinterProcessingStateDetail_TrimmerEmpty,
		"trimmerfull":                               PrinterProcessingStateDetail_TrimmerFull,
		"trimmerinterlockclosed":                    PrinterProcessingStateDetail_TrimmerInterlockClosed,
		"trimmerinterlockopen":                      PrinterProcessingStateDetail_TrimmerInterlockOpen,
		"trimmerjam":                                PrinterProcessingStateDetail_TrimmerJam,
		"trimmerlifealmostover":                     PrinterProcessingStateDetail_TrimmerLifeAlmostOver,
		"trimmerlifeover":                           PrinterProcessingStateDetail_TrimmerLifeOver,
		"trimmermemoryexhausted":                    PrinterProcessingStateDetail_TrimmerMemoryExhausted,
		"trimmermissing":                            PrinterProcessingStateDetail_TrimmerMissing,
		"trimmermotorfailure":                       PrinterProcessingStateDetail_TrimmerMotorFailure,
		"trimmernearlimit":                          PrinterProcessingStateDetail_TrimmerNearLimit,
		"trimmeroffline":                            PrinterProcessingStateDetail_TrimmerOffline,
		"trimmeropened":                             PrinterProcessingStateDetail_TrimmerOpened,
		"trimmerovertemperature":                    PrinterProcessingStateDetail_TrimmerOverTemperature,
		"trimmerpowersaver":                         PrinterProcessingStateDetail_TrimmerPowerSaver,
		"trimmerrecoverablefailure":                 PrinterProcessingStateDetail_TrimmerRecoverableFailure,
		"trimmerrecoverablestorage":                 PrinterProcessingStateDetail_TrimmerRecoverableStorage,
		"trimmerremoved":                            PrinterProcessingStateDetail_TrimmerRemoved,
		"trimmerresourceadded":                      PrinterProcessingStateDetail_TrimmerResourceAdded,
		"trimmerresourceremoved":                    PrinterProcessingStateDetail_TrimmerResourceRemoved,
		"trimmerthermistorfailure":                  PrinterProcessingStateDetail_TrimmerThermistorFailure,
		"trimmertimingfailure":                      PrinterProcessingStateDetail_TrimmerTimingFailure,
		"trimmerturnedoff":                          PrinterProcessingStateDetail_TrimmerTurnedOff,
		"trimmerturnedon":                           PrinterProcessingStateDetail_TrimmerTurnedOn,
		"trimmerundertemperature":                   PrinterProcessingStateDetail_TrimmerUnderTemperature,
		"trimmerunrecoverablefailure":               PrinterProcessingStateDetail_TrimmerUnrecoverableFailure,
		"trimmerunrecoverablestorageerror":          PrinterProcessingStateDetail_TrimmerUnrecoverableStorageError,
		"trimmerwarmingup":                          PrinterProcessingStateDetail_TrimmerWarmingUp,
		"unknown":                                   PrinterProcessingStateDetail_Unknown,
		"wrapperadded":                              PrinterProcessingStateDetail_WrapperAdded,
		"wrapperalmostempty":                        PrinterProcessingStateDetail_WrapperAlmostEmpty,
		"wrapperalmostfull":                         PrinterProcessingStateDetail_WrapperAlmostFull,
		"wrapperatlimit":                            PrinterProcessingStateDetail_WrapperAtLimit,
		"wrapperclosed":                             PrinterProcessingStateDetail_WrapperClosed,
		"wrapperconfigurationchange":                PrinterProcessingStateDetail_WrapperConfigurationChange,
		"wrappercoverclosed":                        PrinterProcessingStateDetail_WrapperCoverClosed,
		"wrappercoveropen":                          PrinterProcessingStateDetail_WrapperCoverOpen,
		"wrapperempty":                              PrinterProcessingStateDetail_WrapperEmpty,
		"wrapperfull":                               PrinterProcessingStateDetail_WrapperFull,
		"wrapperinterlockclosed":                    PrinterProcessingStateDetail_WrapperInterlockClosed,
		"wrapperinterlockopen":                      PrinterProcessingStateDetail_WrapperInterlockOpen,
		"wrapperjam":                                PrinterProcessingStateDetail_WrapperJam,
		"wrapperlifealmostover":                     PrinterProcessingStateDetail_WrapperLifeAlmostOver,
		"wrapperlifeover":                           PrinterProcessingStateDetail_WrapperLifeOver,
		"wrappermemoryexhausted":                    PrinterProcessingStateDetail_WrapperMemoryExhausted,
		"wrappermissing":                            PrinterProcessingStateDetail_WrapperMissing,
		"wrappermotorfailure":                       PrinterProcessingStateDetail_WrapperMotorFailure,
		"wrappernearlimit":                          PrinterProcessingStateDetail_WrapperNearLimit,
		"wrapperoffline":                            PrinterProcessingStateDetail_WrapperOffline,
		"wrapperopened":                             PrinterProcessingStateDetail_WrapperOpened,
		"wrapperovertemperature":                    PrinterProcessingStateDetail_WrapperOverTemperature,
		"wrapperpowersaver":                         PrinterProcessingStateDetail_WrapperPowerSaver,
		"wrapperrecoverablefailure":                 PrinterProcessingStateDetail_WrapperRecoverableFailure,
		"wrapperrecoverablestorage":                 PrinterProcessingStateDetail_WrapperRecoverableStorage,
		"wrapperremoved":                            PrinterProcessingStateDetail_WrapperRemoved,
		"wrapperresourceadded":                      PrinterProcessingStateDetail_WrapperResourceAdded,
		"wrapperresourceremoved":                    PrinterProcessingStateDetail_WrapperResourceRemoved,
		"wrapperthermistorfailure":                  PrinterProcessingStateDetail_WrapperThermistorFailure,
		"wrappertimingfailure":                      PrinterProcessingStateDetail_WrapperTimingFailure,
		"wrapperturnedoff":                          PrinterProcessingStateDetail_WrapperTurnedOff,
		"wrapperturnedon":                           PrinterProcessingStateDetail_WrapperTurnedOn,
		"wrapperundertemperature":                   PrinterProcessingStateDetail_WrapperUnderTemperature,
		"wrapperunrecoverablefailure":               PrinterProcessingStateDetail_WrapperUnrecoverableFailure,
		"wrapperunrecoverablestorageerror":          PrinterProcessingStateDetail_WrapperUnrecoverableStorageError,
		"wrapperwarmingup":                          PrinterProcessingStateDetail_WrapperWarmingUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterProcessingStateDetail(input)
	return &out, nil
}
