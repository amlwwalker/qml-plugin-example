

#pragma once

#ifndef GO_MOC_5292b1_H
#define GO_MOC_5292b1_H

#include <stdint.h>

#ifdef __cplusplus
class Controller5292b1;
void Controller5292b1_Controller5292b1_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
int Controller5292b1_Controller5292b1_QRegisterMetaType();
int Controller5292b1_Controller5292b1_QRegisterMetaType2(char* typeName);
int Controller5292b1_Controller5292b1_QmlRegisterType();
int Controller5292b1_Controller5292b1_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* Controller5292b1___children_atList(void* ptr, int i);
void Controller5292b1___children_setList(void* ptr, void* i);
void* Controller5292b1___children_newList(void* ptr);
void* Controller5292b1___dynamicPropertyNames_atList(void* ptr, int i);
void Controller5292b1___dynamicPropertyNames_setList(void* ptr, void* i);
void* Controller5292b1___dynamicPropertyNames_newList(void* ptr);
void* Controller5292b1___findChildren_atList(void* ptr, int i);
void Controller5292b1___findChildren_setList(void* ptr, void* i);
void* Controller5292b1___findChildren_newList(void* ptr);
void* Controller5292b1___findChildren_atList3(void* ptr, int i);
void Controller5292b1___findChildren_setList3(void* ptr, void* i);
void* Controller5292b1___findChildren_newList3(void* ptr);
void* Controller5292b1___qFindChildren_atList2(void* ptr, int i);
void Controller5292b1___qFindChildren_setList2(void* ptr, void* i);
void* Controller5292b1___qFindChildren_newList2(void* ptr);
void* Controller5292b1_NewController(void* parent);
void Controller5292b1_DestroyController(void* ptr);
void Controller5292b1_DestroyControllerDefault(void* ptr);
void Controller5292b1_ChildEventDefault(void* ptr, void* event);
void Controller5292b1_ConnectNotifyDefault(void* ptr, void* sign);
void Controller5292b1_CustomEventDefault(void* ptr, void* event);
void Controller5292b1_DeleteLaterDefault(void* ptr);
void Controller5292b1_DisconnectNotifyDefault(void* ptr, void* sign);
char Controller5292b1_EventDefault(void* ptr, void* e);
char Controller5292b1_EventFilterDefault(void* ptr, void* watched, void* event);
;
void Controller5292b1_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif