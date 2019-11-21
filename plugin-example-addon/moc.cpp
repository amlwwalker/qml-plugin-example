

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class Controller5292b1: public QObject
{
Q_OBJECT
public:
	Controller5292b1(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Controller5292b1_Controller5292b1_QRegisterMetaType();Controller5292b1_Controller5292b1_QRegisterMetaTypes();callbackController5292b1_Constructor(this);};
	 ~Controller5292b1() { callbackController5292b1_DestroyController(this); };
	void childEvent(QChildEvent * event) { callbackController5292b1_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackController5292b1_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackController5292b1_CustomEvent(this, event); };
	void deleteLater() { callbackController5292b1_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackController5292b1_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackController5292b1_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackController5292b1_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackController5292b1_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackController5292b1_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackController5292b1_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(Controller5292b1*)


void Controller5292b1_Controller5292b1_QRegisterMetaTypes() {
}

int Controller5292b1_Controller5292b1_QRegisterMetaType()
{
	return qRegisterMetaType<Controller5292b1*>();
}

int Controller5292b1_Controller5292b1_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Controller5292b1*>(const_cast<const char*>(typeName));
}

int Controller5292b1_Controller5292b1_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Controller5292b1>();
#else
	return 0;
#endif
}

int Controller5292b1_Controller5292b1_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Controller5292b1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Controller5292b1___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Controller5292b1___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Controller5292b1___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Controller5292b1___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Controller5292b1___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Controller5292b1___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Controller5292b1___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Controller5292b1___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Controller5292b1___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Controller5292b1___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Controller5292b1___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Controller5292b1___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Controller5292b1___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Controller5292b1___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Controller5292b1___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Controller5292b1_NewController(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Controller5292b1(static_cast<QWindow*>(parent));
	} else {
		return new Controller5292b1(static_cast<QObject*>(parent));
	}
}

void Controller5292b1_DestroyController(void* ptr)
{
	static_cast<Controller5292b1*>(ptr)->~Controller5292b1();
}

void Controller5292b1_DestroyControllerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void Controller5292b1_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Controller5292b1*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Controller5292b1_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Controller5292b1*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Controller5292b1_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Controller5292b1*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Controller5292b1_DeleteLaterDefault(void* ptr)
{
	static_cast<Controller5292b1*>(ptr)->QObject::deleteLater();
}

void Controller5292b1_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Controller5292b1*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char Controller5292b1_EventDefault(void* ptr, void* e)
{
	return static_cast<Controller5292b1*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Controller5292b1_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Controller5292b1*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void Controller5292b1_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Controller5292b1*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
