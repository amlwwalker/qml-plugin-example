

#pragma once

#ifndef GO_MOC_283f7e_H
#define GO_MOC_283f7e_H

#include <stdint.h>

#ifdef __cplusplus
class Controller283f7e;
void Controller283f7e_Controller283f7e_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
int Controller283f7e_Controller283f7e_QRegisterMetaType();
int Controller283f7e_Controller283f7e_QRegisterMetaType2(char* typeName);
int Controller283f7e_Controller283f7e_QmlRegisterType();
int Controller283f7e_Controller283f7e_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* Controller283f7e___children_atList(void* ptr, int i);
void Controller283f7e___children_setList(void* ptr, void* i);
void* Controller283f7e___children_newList(void* ptr);
void* Controller283f7e___dynamicPropertyNames_atList(void* ptr, int i);
void Controller283f7e___dynamicPropertyNames_setList(void* ptr, void* i);
void* Controller283f7e___dynamicPropertyNames_newList(void* ptr);
void* Controller283f7e___findChildren_atList(void* ptr, int i);
void Controller283f7e___findChildren_setList(void* ptr, void* i);
void* Controller283f7e___findChildren_newList(void* ptr);
void* Controller283f7e___findChildren_atList3(void* ptr, int i);
void Controller283f7e___findChildren_setList3(void* ptr, void* i);
void* Controller283f7e___findChildren_newList3(void* ptr);
void* Controller283f7e___qFindChildren_atList2(void* ptr, int i);
void Controller283f7e___qFindChildren_setList2(void* ptr, void* i);
void* Controller283f7e___qFindChildren_newList2(void* ptr);
void* Controller283f7e_NewController(void* parent);
void Controller283f7e_DestroyController(void* ptr);
void Controller283f7e_DestroyControllerDefault(void* ptr);
void Controller283f7e_ChildEventDefault(void* ptr, void* event);
void Controller283f7e_ConnectNotifyDefault(void* ptr, void* sign);
void Controller283f7e_CustomEventDefault(void* ptr, void* event);
void Controller283f7e_DeleteLaterDefault(void* ptr);
void Controller283f7e_DisconnectNotifyDefault(void* ptr, void* sign);
char Controller283f7e_EventDefault(void* ptr, void* e);
char Controller283f7e_EventFilterDefault(void* ptr, void* watched, void* event);
;
void Controller283f7e_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif