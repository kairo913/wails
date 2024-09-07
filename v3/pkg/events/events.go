package events

type ApplicationEventType uint
type WindowEventType uint

var Common = newCommonEvents()

type commonEvents struct {
	ApplicationStarted ApplicationEventType
	WindowMaximise     WindowEventType
	WindowUnMaximise   WindowEventType
	WindowFullscreen   WindowEventType
	WindowUnFullscreen WindowEventType
	WindowRestore      WindowEventType
	WindowMinimise     WindowEventType
	WindowUnMinimise   WindowEventType
	WindowClosing      WindowEventType
	WindowZoom         WindowEventType
	WindowZoomIn       WindowEventType
	WindowZoomOut      WindowEventType
	WindowZoomReset    WindowEventType
	WindowFocus        WindowEventType
	WindowLostFocus    WindowEventType
	WindowShow         WindowEventType
	WindowHide         WindowEventType
	WindowDPIChanged   WindowEventType
	WindowFilesDropped WindowEventType
	WindowRuntimeReady WindowEventType
	ThemeChanged       ApplicationEventType
	WindowDidMove      WindowEventType
	WindowDidResize    WindowEventType
	WindowCreated      WindowEventType
	WindowLoaded       WindowEventType
}

func newCommonEvents() commonEvents {
	return commonEvents{
		ApplicationStarted: 1183,
		WindowMaximise:     1184,
		WindowUnMaximise:   1185,
		WindowFullscreen:   1186,
		WindowUnFullscreen: 1187,
		WindowRestore:      1188,
		WindowMinimise:     1189,
		WindowUnMinimise:   1190,
		WindowClosing:      1191,
		WindowZoom:         1192,
		WindowZoomIn:       1193,
		WindowZoomOut:      1194,
		WindowZoomReset:    1195,
		WindowFocus:        1196,
		WindowLostFocus:    1197,
		WindowShow:         1198,
		WindowHide:         1199,
		WindowDPIChanged:   1200,
		WindowFilesDropped: 1201,
		WindowRuntimeReady: 1202,
		ThemeChanged:       1203,
		WindowDidMove:      1204,
		WindowDidResize:    1205,
		WindowCreated:      1206,
		WindowLoaded:       1207,
	}
}

var Linux = newLinuxEvents()

type linuxEvents struct {
	SystemThemeChanged ApplicationEventType
	WindowLoadChanged  WindowEventType
	WindowDeleteEvent  WindowEventType
	WindowDidMove      WindowEventType
	WindowDidResize    WindowEventType
	WindowFocusIn      WindowEventType
	WindowFocusOut     WindowEventType
	ApplicationStartup ApplicationEventType
}

func newLinuxEvents() linuxEvents {
	return linuxEvents{
		SystemThemeChanged: 1024,
		WindowLoadChanged:  1025,
		WindowDeleteEvent:  1026,
		WindowDidMove:      1027,
		WindowDidResize:    1028,
		WindowFocusIn:      1029,
		WindowFocusOut:     1030,
		ApplicationStartup: 1031,
	}
}

var Mac = newMacEvents()

type macEvents struct {
	ApplicationDidBecomeActive                              ApplicationEventType
	ApplicationDidChangeBackingProperties                   ApplicationEventType
	ApplicationDidChangeEffectiveAppearance                 ApplicationEventType
	ApplicationDidChangeIcon                                ApplicationEventType
	ApplicationDidChangeOcclusionState                      ApplicationEventType
	ApplicationDidChangeScreenParameters                    ApplicationEventType
	ApplicationDidChangeStatusBarFrame                      ApplicationEventType
	ApplicationDidChangeStatusBarOrientation                ApplicationEventType
	ApplicationDidFinishLaunching                           ApplicationEventType
	ApplicationDidHide                                      ApplicationEventType
	ApplicationDidResignActiveNotification                  ApplicationEventType
	ApplicationDidUnhide                                    ApplicationEventType
	ApplicationDidUpdate                                    ApplicationEventType
	ApplicationWillBecomeActive                             ApplicationEventType
	ApplicationWillFinishLaunching                          ApplicationEventType
	ApplicationWillHide                                     ApplicationEventType
	ApplicationWillResignActive                             ApplicationEventType
	ApplicationWillTerminate                                ApplicationEventType
	ApplicationWillUnhide                                   ApplicationEventType
	ApplicationWillUpdate                                   ApplicationEventType
	ApplicationDidChangeTheme                               ApplicationEventType
	ApplicationShouldHandleReopen                           ApplicationEventType
	WindowDidBecomeKey                                      WindowEventType
	WindowDidBecomeMain                                     WindowEventType
	WindowDidBeginSheet                                     WindowEventType
	WindowDidChangeAlpha                                    WindowEventType
	WindowDidChangeBackingLocation                          WindowEventType
	WindowDidChangeBackingProperties                        WindowEventType
	WindowDidChangeCollectionBehavior                       WindowEventType
	WindowDidChangeEffectiveAppearance                      WindowEventType
	WindowDidChangeOcclusionState                           WindowEventType
	WindowDidChangeOrderingMode                             WindowEventType
	WindowDidChangeScreen                                   WindowEventType
	WindowDidChangeScreenParameters                         WindowEventType
	WindowDidChangeScreenProfile                            WindowEventType
	WindowDidChangeScreenSpace                              WindowEventType
	WindowDidChangeScreenSpaceProperties                    WindowEventType
	WindowDidChangeSharingType                              WindowEventType
	WindowDidChangeSpace                                    WindowEventType
	WindowDidChangeSpaceOrderingMode                        WindowEventType
	WindowDidChangeTitle                                    WindowEventType
	WindowDidChangeToolbar                                  WindowEventType
	WindowDidChangeVisibility                               WindowEventType
	WindowDidDeminiaturize                                  WindowEventType
	WindowDidEndSheet                                       WindowEventType
	WindowDidEnterFullScreen                                WindowEventType
	WindowDidEnterVersionBrowser                            WindowEventType
	WindowDidExitFullScreen                                 WindowEventType
	WindowDidExitVersionBrowser                             WindowEventType
	WindowDidExpose                                         WindowEventType
	WindowDidFocus                                          WindowEventType
	WindowDidMiniaturize                                    WindowEventType
	WindowDidMove                                           WindowEventType
	WindowDidOrderOffScreen                                 WindowEventType
	WindowDidOrderOnScreen                                  WindowEventType
	WindowDidResignKey                                      WindowEventType
	WindowDidResignMain                                     WindowEventType
	WindowDidResize                                         WindowEventType
	WindowDidUpdate                                         WindowEventType
	WindowDidUpdateAlpha                                    WindowEventType
	WindowDidUpdateCollectionBehavior                       WindowEventType
	WindowDidUpdateCollectionProperties                     WindowEventType
	WindowDidUpdateShadow                                   WindowEventType
	WindowDidUpdateTitle                                    WindowEventType
	WindowDidUpdateToolbar                                  WindowEventType
	WindowDidUpdateVisibility                               WindowEventType
	WindowShouldClose                                       WindowEventType
	WindowWillBecomeKey                                     WindowEventType
	WindowWillBecomeMain                                    WindowEventType
	WindowWillBeginSheet                                    WindowEventType
	WindowWillChangeOrderingMode                            WindowEventType
	WindowWillClose                                         WindowEventType
	WindowWillDeminiaturize                                 WindowEventType
	WindowWillEnterFullScreen                               WindowEventType
	WindowWillEnterVersionBrowser                           WindowEventType
	WindowWillExitFullScreen                                WindowEventType
	WindowWillExitVersionBrowser                            WindowEventType
	WindowWillFocus                                         WindowEventType
	WindowWillMiniaturize                                   WindowEventType
	WindowWillMove                                          WindowEventType
	WindowWillOrderOffScreen                                WindowEventType
	WindowWillOrderOnScreen                                 WindowEventType
	WindowWillResignMain                                    WindowEventType
	WindowWillResize                                        WindowEventType
	WindowWillUnfocus                                       WindowEventType
	WindowWillUpdate                                        WindowEventType
	WindowWillUpdateAlpha                                   WindowEventType
	WindowWillUpdateCollectionBehavior                      WindowEventType
	WindowWillUpdateCollectionProperties                    WindowEventType
	WindowWillUpdateShadow                                  WindowEventType
	WindowWillUpdateTitle                                   WindowEventType
	WindowWillUpdateToolbar                                 WindowEventType
	WindowWillUpdateVisibility                              WindowEventType
	WindowWillUseStandardFrame                              WindowEventType
	MenuWillOpen                                            ApplicationEventType
	MenuDidOpen                                             ApplicationEventType
	MenuDidClose                                            ApplicationEventType
	MenuWillSendAction                                      ApplicationEventType
	MenuDidSendAction                                       ApplicationEventType
	MenuWillHighlightItem                                   ApplicationEventType
	MenuDidHighlightItem                                    ApplicationEventType
	MenuWillDisplayItem                                     ApplicationEventType
	MenuDidDisplayItem                                      ApplicationEventType
	MenuWillAddItem                                         ApplicationEventType
	MenuDidAddItem                                          ApplicationEventType
	MenuWillRemoveItem                                      ApplicationEventType
	MenuDidRemoveItem                                       ApplicationEventType
	MenuWillBeginTracking                                   ApplicationEventType
	MenuDidBeginTracking                                    ApplicationEventType
	MenuWillEndTracking                                     ApplicationEventType
	MenuDidEndTracking                                      ApplicationEventType
	MenuWillUpdate                                          ApplicationEventType
	MenuDidUpdate                                           ApplicationEventType
	MenuWillPopUp                                           ApplicationEventType
	MenuDidPopUp                                            ApplicationEventType
	MenuWillSendActionToItem                                ApplicationEventType
	MenuDidSendActionToItem                                 ApplicationEventType
	WebViewDidStartProvisionalNavigation                    WindowEventType
	WebViewDidReceiveServerRedirectForProvisionalNavigation WindowEventType
	WebViewDidFinishNavigation                              WindowEventType
	WebViewDidCommitNavigation                              WindowEventType
	WindowFileDraggingEntered                               WindowEventType
	WindowFileDraggingPerformed                             WindowEventType
	WindowFileDraggingExited                                WindowEventType
}

func newMacEvents() macEvents {
	return macEvents{
		ApplicationDidBecomeActive:               1032,
		ApplicationDidChangeBackingProperties:    1033,
		ApplicationDidChangeEffectiveAppearance:  1034,
		ApplicationDidChangeIcon:                 1035,
		ApplicationDidChangeOcclusionState:       1036,
		ApplicationDidChangeScreenParameters:     1037,
		ApplicationDidChangeStatusBarFrame:       1038,
		ApplicationDidChangeStatusBarOrientation: 1039,
		ApplicationDidFinishLaunching:            1040,
		ApplicationDidHide:                       1041,
		ApplicationDidResignActiveNotification:   1042,
		ApplicationDidUnhide:                     1043,
		ApplicationDidUpdate:                     1044,
		ApplicationWillBecomeActive:              1045,
		ApplicationWillFinishLaunching:           1046,
		ApplicationWillHide:                      1047,
		ApplicationWillResignActive:              1048,
		ApplicationWillTerminate:                 1049,
		ApplicationWillUnhide:                    1050,
		ApplicationWillUpdate:                    1051,
		ApplicationDidChangeTheme:                1052,
		ApplicationShouldHandleReopen:            1053,
		WindowDidBecomeKey:                       1054,
		WindowDidBecomeMain:                      1055,
		WindowDidBeginSheet:                      1056,
		WindowDidChangeAlpha:                     1057,
		WindowDidChangeBackingLocation:           1058,
		WindowDidChangeBackingProperties:         1059,
		WindowDidChangeCollectionBehavior:        1060,
		WindowDidChangeEffectiveAppearance:       1061,
		WindowDidChangeOcclusionState:            1062,
		WindowDidChangeOrderingMode:              1063,
		WindowDidChangeScreen:                    1064,
		WindowDidChangeScreenParameters:          1065,
		WindowDidChangeScreenProfile:             1066,
		WindowDidChangeScreenSpace:               1067,
		WindowDidChangeScreenSpaceProperties:     1068,
		WindowDidChangeSharingType:               1069,
		WindowDidChangeSpace:                     1070,
		WindowDidChangeSpaceOrderingMode:         1071,
		WindowDidChangeTitle:                     1072,
		WindowDidChangeToolbar:                   1073,
		WindowDidChangeVisibility:                1074,
		WindowDidDeminiaturize:                   1075,
		WindowDidEndSheet:                        1076,
		WindowDidEnterFullScreen:                 1077,
		WindowDidEnterVersionBrowser:             1078,
		WindowDidExitFullScreen:                  1079,
		WindowDidExitVersionBrowser:              1080,
		WindowDidExpose:                          1081,
		WindowDidFocus:                           1082,
		WindowDidMiniaturize:                     1083,
		WindowDidMove:                            1084,
		WindowDidOrderOffScreen:                  1085,
		WindowDidOrderOnScreen:                   1086,
		WindowDidResignKey:                       1087,
		WindowDidResignMain:                      1088,
		WindowDidResize:                          1089,
		WindowDidUpdate:                          1090,
		WindowDidUpdateAlpha:                     1091,
		WindowDidUpdateCollectionBehavior:        1092,
		WindowDidUpdateCollectionProperties:      1093,
		WindowDidUpdateShadow:                    1094,
		WindowDidUpdateTitle:                     1095,
		WindowDidUpdateToolbar:                   1096,
		WindowDidUpdateVisibility:                1097,
		WindowShouldClose:                        1098,
		WindowWillBecomeKey:                      1099,
		WindowWillBecomeMain:                     1100,
		WindowWillBeginSheet:                     1101,
		WindowWillChangeOrderingMode:             1102,
		WindowWillClose:                          1103,
		WindowWillDeminiaturize:                  1104,
		WindowWillEnterFullScreen:                1105,
		WindowWillEnterVersionBrowser:            1106,
		WindowWillExitFullScreen:                 1107,
		WindowWillExitVersionBrowser:             1108,
		WindowWillFocus:                          1109,
		WindowWillMiniaturize:                    1110,
		WindowWillMove:                           1111,
		WindowWillOrderOffScreen:                 1112,
		WindowWillOrderOnScreen:                  1113,
		WindowWillResignMain:                     1114,
		WindowWillResize:                         1115,
		WindowWillUnfocus:                        1116,
		WindowWillUpdate:                         1117,
		WindowWillUpdateAlpha:                    1118,
		WindowWillUpdateCollectionBehavior:       1119,
		WindowWillUpdateCollectionProperties:     1120,
		WindowWillUpdateShadow:                   1121,
		WindowWillUpdateTitle:                    1122,
		WindowWillUpdateToolbar:                  1123,
		WindowWillUpdateVisibility:               1124,
		WindowWillUseStandardFrame:               1125,
		MenuWillOpen:                             1126,
		MenuDidOpen:                              1127,
		MenuDidClose:                             1128,
		MenuWillSendAction:                       1129,
		MenuDidSendAction:                        1130,
		MenuWillHighlightItem:                    1131,
		MenuDidHighlightItem:                     1132,
		MenuWillDisplayItem:                      1133,
		MenuDidDisplayItem:                       1134,
		MenuWillAddItem:                          1135,
		MenuDidAddItem:                           1136,
		MenuWillRemoveItem:                       1137,
		MenuDidRemoveItem:                        1138,
		MenuWillBeginTracking:                    1139,
		MenuDidBeginTracking:                     1140,
		MenuWillEndTracking:                      1141,
		MenuDidEndTracking:                       1142,
		MenuWillUpdate:                           1143,
		MenuDidUpdate:                            1144,
		MenuWillPopUp:                            1145,
		MenuDidPopUp:                             1146,
		MenuWillSendActionToItem:                 1147,
		MenuDidSendActionToItem:                  1148,
		WebViewDidStartProvisionalNavigation:     1149,
		WebViewDidReceiveServerRedirectForProvisionalNavigation: 1150,
		WebViewDidFinishNavigation:                              1151,
		WebViewDidCommitNavigation:                              1152,
		WindowFileDraggingEntered:                               1153,
		WindowFileDraggingPerformed:                             1154,
		WindowFileDraggingExited:                                1155,
	}
}

var Windows = newWindowsEvents()

type windowsEvents struct {
	SystemThemeChanged         ApplicationEventType
	APMPowerStatusChange       ApplicationEventType
	APMSuspend                 ApplicationEventType
	APMResumeAutomatic         ApplicationEventType
	APMResumeSuspend           ApplicationEventType
	APMPowerSettingChange      ApplicationEventType
	ApplicationStarted         ApplicationEventType
	WebViewNavigationCompleted WindowEventType
	WindowInactive             WindowEventType
	WindowActive               WindowEventType
	WindowClickActive          WindowEventType
	WindowMaximise             WindowEventType
	WindowUnMaximise           WindowEventType
	WindowFullscreen           WindowEventType
	WindowUnFullscreen         WindowEventType
	WindowRestore              WindowEventType
	WindowMinimise             WindowEventType
	WindowUnMinimise           WindowEventType
	WindowClose                WindowEventType
	WindowSetFocus             WindowEventType
	WindowKillFocus            WindowEventType
	WindowDragDrop             WindowEventType
	WindowDragEnter            WindowEventType
	WindowDragLeave            WindowEventType
	WindowDragOver             WindowEventType
	WindowDidMove              WindowEventType
	WindowDidResize            WindowEventType
}

func newWindowsEvents() windowsEvents {
	return windowsEvents{
		SystemThemeChanged:         1156,
		APMPowerStatusChange:       1157,
		APMSuspend:                 1158,
		APMResumeAutomatic:         1159,
		APMResumeSuspend:           1160,
		APMPowerSettingChange:      1161,
		ApplicationStarted:         1162,
		WebViewNavigationCompleted: 1163,
		WindowInactive:             1164,
		WindowActive:               1165,
		WindowClickActive:          1166,
		WindowMaximise:             1167,
		WindowUnMaximise:           1168,
		WindowFullscreen:           1169,
		WindowUnFullscreen:         1170,
		WindowRestore:              1171,
		WindowMinimise:             1172,
		WindowUnMinimise:           1173,
		WindowClose:                1174,
		WindowSetFocus:             1175,
		WindowKillFocus:            1176,
		WindowDragDrop:             1177,
		WindowDragEnter:            1178,
		WindowDragLeave:            1179,
		WindowDragOver:             1180,
		WindowDidMove:              1181,
		WindowDidResize:            1182,
	}
}

func JSEvent(event uint) string {
	return eventToJS[event]
}

var eventToJS = map[uint]string{
	1024: "linux:SystemThemeChanged",
	1025: "linux:WindowLoadChanged",
	1026: "linux:WindowDeleteEvent",
	1027: "linux:WindowDidMove",
	1028: "linux:WindowDidResize",
	1029: "linux:WindowFocusIn",
	1030: "linux:WindowFocusOut",
	1031: "linux:ApplicationStartup",
	1032: "mac:ApplicationDidBecomeActive",
	1033: "mac:ApplicationDidChangeBackingProperties",
	1034: "mac:ApplicationDidChangeEffectiveAppearance",
	1035: "mac:ApplicationDidChangeIcon",
	1036: "mac:ApplicationDidChangeOcclusionState",
	1037: "mac:ApplicationDidChangeScreenParameters",
	1038: "mac:ApplicationDidChangeStatusBarFrame",
	1039: "mac:ApplicationDidChangeStatusBarOrientation",
	1040: "mac:ApplicationDidFinishLaunching",
	1041: "mac:ApplicationDidHide",
	1042: "mac:ApplicationDidResignActiveNotification",
	1043: "mac:ApplicationDidUnhide",
	1044: "mac:ApplicationDidUpdate",
	1045: "mac:ApplicationWillBecomeActive",
	1046: "mac:ApplicationWillFinishLaunching",
	1047: "mac:ApplicationWillHide",
	1048: "mac:ApplicationWillResignActive",
	1049: "mac:ApplicationWillTerminate",
	1050: "mac:ApplicationWillUnhide",
	1051: "mac:ApplicationWillUpdate",
	1052: "mac:ApplicationDidChangeTheme!",
	1053: "mac:ApplicationShouldHandleReopen!",
	1054: "mac:WindowDidBecomeKey",
	1055: "mac:WindowDidBecomeMain",
	1056: "mac:WindowDidBeginSheet",
	1057: "mac:WindowDidChangeAlpha",
	1058: "mac:WindowDidChangeBackingLocation",
	1059: "mac:WindowDidChangeBackingProperties",
	1060: "mac:WindowDidChangeCollectionBehavior",
	1061: "mac:WindowDidChangeEffectiveAppearance",
	1062: "mac:WindowDidChangeOcclusionState",
	1063: "mac:WindowDidChangeOrderingMode",
	1064: "mac:WindowDidChangeScreen",
	1065: "mac:WindowDidChangeScreenParameters",
	1066: "mac:WindowDidChangeScreenProfile",
	1067: "mac:WindowDidChangeScreenSpace",
	1068: "mac:WindowDidChangeScreenSpaceProperties",
	1069: "mac:WindowDidChangeSharingType",
	1070: "mac:WindowDidChangeSpace",
	1071: "mac:WindowDidChangeSpaceOrderingMode",
	1072: "mac:WindowDidChangeTitle",
	1073: "mac:WindowDidChangeToolbar",
	1074: "mac:WindowDidChangeVisibility",
	1075: "mac:WindowDidDeminiaturize",
	1076: "mac:WindowDidEndSheet",
	1077: "mac:WindowDidEnterFullScreen",
	1078: "mac:WindowDidEnterVersionBrowser",
	1079: "mac:WindowDidExitFullScreen",
	1080: "mac:WindowDidExitVersionBrowser",
	1081: "mac:WindowDidExpose",
	1082: "mac:WindowDidFocus",
	1083: "mac:WindowDidMiniaturize",
	1084: "mac:WindowDidMove",
	1085: "mac:WindowDidOrderOffScreen",
	1086: "mac:WindowDidOrderOnScreen",
	1087: "mac:WindowDidResignKey",
	1088: "mac:WindowDidResignMain",
	1089: "mac:WindowDidResize",
	1090: "mac:WindowDidUpdate",
	1091: "mac:WindowDidUpdateAlpha",
	1092: "mac:WindowDidUpdateCollectionBehavior",
	1093: "mac:WindowDidUpdateCollectionProperties",
	1094: "mac:WindowDidUpdateShadow",
	1095: "mac:WindowDidUpdateTitle",
	1096: "mac:WindowDidUpdateToolbar",
	1097: "mac:WindowDidUpdateVisibility",
	1098: "mac:WindowShouldClose!",
	1099: "mac:WindowWillBecomeKey",
	1100: "mac:WindowWillBecomeMain",
	1101: "mac:WindowWillBeginSheet",
	1102: "mac:WindowWillChangeOrderingMode",
	1103: "mac:WindowWillClose",
	1104: "mac:WindowWillDeminiaturize",
	1105: "mac:WindowWillEnterFullScreen",
	1106: "mac:WindowWillEnterVersionBrowser",
	1107: "mac:WindowWillExitFullScreen",
	1108: "mac:WindowWillExitVersionBrowser",
	1109: "mac:WindowWillFocus",
	1110: "mac:WindowWillMiniaturize",
	1111: "mac:WindowWillMove",
	1112: "mac:WindowWillOrderOffScreen",
	1113: "mac:WindowWillOrderOnScreen",
	1114: "mac:WindowWillResignMain",
	1115: "mac:WindowWillResize",
	1116: "mac:WindowWillUnfocus",
	1117: "mac:WindowWillUpdate",
	1118: "mac:WindowWillUpdateAlpha",
	1119: "mac:WindowWillUpdateCollectionBehavior",
	1120: "mac:WindowWillUpdateCollectionProperties",
	1121: "mac:WindowWillUpdateShadow",
	1122: "mac:WindowWillUpdateTitle",
	1123: "mac:WindowWillUpdateToolbar",
	1124: "mac:WindowWillUpdateVisibility",
	1125: "mac:WindowWillUseStandardFrame",
	1126: "mac:MenuWillOpen",
	1127: "mac:MenuDidOpen",
	1128: "mac:MenuDidClose",
	1129: "mac:MenuWillSendAction",
	1130: "mac:MenuDidSendAction",
	1131: "mac:MenuWillHighlightItem",
	1132: "mac:MenuDidHighlightItem",
	1133: "mac:MenuWillDisplayItem",
	1134: "mac:MenuDidDisplayItem",
	1135: "mac:MenuWillAddItem",
	1136: "mac:MenuDidAddItem",
	1137: "mac:MenuWillRemoveItem",
	1138: "mac:MenuDidRemoveItem",
	1139: "mac:MenuWillBeginTracking",
	1140: "mac:MenuDidBeginTracking",
	1141: "mac:MenuWillEndTracking",
	1142: "mac:MenuDidEndTracking",
	1143: "mac:MenuWillUpdate",
	1144: "mac:MenuDidUpdate",
	1145: "mac:MenuWillPopUp",
	1146: "mac:MenuDidPopUp",
	1147: "mac:MenuWillSendActionToItem",
	1148: "mac:MenuDidSendActionToItem",
	1149: "mac:WebViewDidStartProvisionalNavigation",
	1150: "mac:WebViewDidReceiveServerRedirectForProvisionalNavigation",
	1151: "mac:WebViewDidFinishNavigation",
	1152: "mac:WebViewDidCommitNavigation",
	1153: "mac:WindowFileDraggingEntered",
	1154: "mac:WindowFileDraggingPerformed",
	1155: "mac:WindowFileDraggingExited",
	1156: "windows:SystemThemeChanged",
	1157: "windows:APMPowerStatusChange",
	1158: "windows:APMSuspend",
	1159: "windows:APMResumeAutomatic",
	1160: "windows:APMResumeSuspend",
	1161: "windows:APMPowerSettingChange",
	1162: "windows:ApplicationStarted",
	1163: "windows:WebViewNavigationCompleted",
	1164: "windows:WindowInactive",
	1165: "windows:WindowActive",
	1166: "windows:WindowClickActive",
	1167: "windows:WindowMaximise",
	1168: "windows:WindowUnMaximise",
	1169: "windows:WindowFullscreen",
	1170: "windows:WindowUnFullscreen",
	1171: "windows:WindowRestore",
	1172: "windows:WindowMinimise",
	1173: "windows:WindowUnMinimise",
	1174: "windows:WindowClose",
	1175: "windows:WindowSetFocus",
	1176: "windows:WindowKillFocus",
	1177: "windows:WindowDragDrop",
	1178: "windows:WindowDragEnter",
	1179: "windows:WindowDragLeave",
	1180: "windows:WindowDragOver",
	1181: "windows:WindowDidMove",
	1182: "windows:WindowDidResize",
	1183: "common:ApplicationStarted",
	1184: "common:WindowMaximise",
	1185: "common:WindowUnMaximise",
	1186: "common:WindowFullscreen",
	1187: "common:WindowUnFullscreen",
	1188: "common:WindowRestore",
	1189: "common:WindowMinimise",
	1190: "common:WindowUnMinimise",
	1191: "common:WindowClosing",
	1192: "common:WindowZoom",
	1193: "common:WindowZoomIn",
	1194: "common:WindowZoomOut",
	1195: "common:WindowZoomReset",
	1196: "common:WindowFocus",
	1197: "common:WindowLostFocus",
	1198: "common:WindowShow",
	1199: "common:WindowHide",
	1200: "common:WindowDPIChanged",
	1201: "common:WindowFilesDropped",
	1202: "common:WindowRuntimeReady",
	1203: "common:ThemeChanged",
	1204: "common:WindowDidMove",
	1205: "common:WindowDidResize",
	1206: "common:WindowCreated",
	1207: "common:WindowLoaded",
}
