

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHoverEvent>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMouseEvent>
#include <QObject>
#include <QPointF>
#include <QQuickItem>
#include <QQuickWindow>
#include <QRectF>
#include <QSGTextureProvider>
#include <QString>
#include <QTimerEvent>
#include <QTouchEvent>
#include <QVariant>
#include <QVector>
#include <QWheelEvent>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class settingsController15bc46: public QQuickItem
{
Q_OBJECT
public:
	settingsController15bc46(QQuickItem *parent = Q_NULLPTR) : QQuickItem(parent) {qRegisterMetaType<quintptr>("quintptr");settingsController15bc46_settingsController15bc46_QRegisterMetaType();settingsController15bc46_settingsController15bc46_QRegisterMetaTypes();callbacksettingsController15bc46_Constructor(this);};
	void Signal_UpdateSettings(QString id, QString name, QString description) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };QByteArray tcb3291 = description.toUtf8(); Moc_PackedString descriptionPacked = { const_cast<char*>(tcb3291.prepend("WHITESPACE").constData()+10), tcb3291.size()-10 };callbacksettingsController15bc46_UpdateSettings(this, idPacked, namePacked, descriptionPacked); };
	 ~settingsController15bc46() { callbacksettingsController15bc46_DestroySettingsController(this); };
	void Signal_ActiveFocusChanged(bool vbo) { callbacksettingsController15bc46_ActiveFocusChanged(this, vbo); };
	void Signal_ActiveFocusOnTabChanged(bool vbo) { callbacksettingsController15bc46_ActiveFocusOnTabChanged(this, vbo); };
	void Signal_AntialiasingChanged(bool vbo) { callbacksettingsController15bc46_AntialiasingChanged(this, vbo); };
	void Signal_BaselineOffsetChanged(qreal vqr) { callbacksettingsController15bc46_BaselineOffsetChanged(this, vqr); };
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbacksettingsController15bc46_ChildMouseEventFilter(this, item, event) != 0; };
	void Signal_ChildrenRectChanged(const QRectF & vqr) { callbacksettingsController15bc46_ChildrenRectChanged(this, const_cast<QRectF*>(&vqr)); };
	void classBegin() { callbacksettingsController15bc46_ClassBegin(this); };
	void Signal_ClipChanged(bool vbo) { callbacksettingsController15bc46_ClipChanged(this, vbo); };
	void componentComplete() { callbacksettingsController15bc46_ComponentComplete(this); };
	void Signal_ContainmentMaskChanged() { callbacksettingsController15bc46_ContainmentMaskChanged(this); };
	bool contains(const QPointF & point) const { return callbacksettingsController15bc46_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	void dragEnterEvent(QDragEnterEvent * event) { callbacksettingsController15bc46_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbacksettingsController15bc46_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbacksettingsController15bc46_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbacksettingsController15bc46_DropEvent(this, event); };
	void Signal_EnabledChanged() { callbacksettingsController15bc46_EnabledChanged(this); };
	bool event(QEvent * ev) { return callbacksettingsController15bc46_Event(this, ev) != 0; };
	void Signal_FocusChanged(bool vbo) { callbacksettingsController15bc46_FocusChanged(this, vbo); };
	void focusInEvent(QFocusEvent * vqf) { callbacksettingsController15bc46_FocusInEvent(this, vqf); };
	void focusOutEvent(QFocusEvent * vqf) { callbacksettingsController15bc46_FocusOutEvent(this, vqf); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbacksettingsController15bc46_GeometryChanged(this, const_cast<QRectF*>(&newGeometry), const_cast<QRectF*>(&oldGeometry)); };
	void Signal_HeightChanged() { callbacksettingsController15bc46_HeightChanged(this); };
	void hoverEnterEvent(QHoverEvent * event) { callbacksettingsController15bc46_HoverEnterEvent(this, event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbacksettingsController15bc46_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QHoverEvent * event) { callbacksettingsController15bc46_HoverMoveEvent(this, event); };
	void Signal_ImplicitHeightChanged() { callbacksettingsController15bc46_ImplicitHeightChanged(this); };
	void Signal_ImplicitWidthChanged() { callbacksettingsController15bc46_ImplicitWidthChanged(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbacksettingsController15bc46_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbacksettingsController15bc46_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool isTextureProvider() const { return callbacksettingsController15bc46_IsTextureProvider(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void keyPressEvent(QKeyEvent * event) { callbacksettingsController15bc46_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbacksettingsController15bc46_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbacksettingsController15bc46_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbacksettingsController15bc46_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbacksettingsController15bc46_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbacksettingsController15bc46_MouseReleaseEvent(this, event); };
	void mouseUngrabEvent() { callbacksettingsController15bc46_MouseUngrabEvent(this); };
	void Signal_OpacityChanged() { callbacksettingsController15bc46_OpacityChanged(this); };
	void Signal_ParentChanged(QQuickItem * vqq) { callbacksettingsController15bc46_ParentChanged(this, vqq); };
	void releaseResources() { callbacksettingsController15bc46_ReleaseResources(this); };
	void Signal_RotationChanged() { callbacksettingsController15bc46_RotationChanged(this); };
	void Signal_ScaleChanged() { callbacksettingsController15bc46_ScaleChanged(this); };
	void Signal_SmoothChanged(bool vbo) { callbacksettingsController15bc46_SmoothChanged(this, vbo); };
	void Signal_StateChanged(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbacksettingsController15bc46_StateChanged(this, vqsPacked); };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbacksettingsController15bc46_TextureProvider(const_cast<void*>(static_cast<const void*>(this)))); };
	void touchEvent(QTouchEvent * event) { callbacksettingsController15bc46_TouchEvent(this, event); };
	void touchUngrabEvent() { callbacksettingsController15bc46_TouchUngrabEvent(this); };
	void Signal_TransformOriginChanged(QQuickItem::TransformOrigin vqq) { callbacksettingsController15bc46_TransformOriginChanged(this, vqq); };
	void update() { callbacksettingsController15bc46_Update(this); };
	void updatePolish() { callbacksettingsController15bc46_UpdatePolish(this); };
	void Signal_VisibleChanged() { callbacksettingsController15bc46_VisibleChanged(this); };
	void wheelEvent(QWheelEvent * event) { callbacksettingsController15bc46_WheelEvent(this, event); };
	void Signal_WidthChanged() { callbacksettingsController15bc46_WidthChanged(this); };
	void Signal_WindowChanged(QQuickWindow * window) { callbacksettingsController15bc46_WindowChanged(this, window); };
	void Signal_XChanged() { callbacksettingsController15bc46_XChanged(this); };
	void Signal_YChanged() { callbacksettingsController15bc46_YChanged(this); };
	void Signal_ZChanged() { callbacksettingsController15bc46_ZChanged(this); };
	void childEvent(QChildEvent * event) { callbacksettingsController15bc46_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbacksettingsController15bc46_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbacksettingsController15bc46_CustomEvent(this, event); };
	void deleteLater() { callbacksettingsController15bc46_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbacksettingsController15bc46_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbacksettingsController15bc46_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbacksettingsController15bc46_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbacksettingsController15bc46_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbacksettingsController15bc46_TimerEvent(this, event); };
signals:
	void updateSettings(QString id, QString name, QString description);
public slots:
private:
};

Q_DECLARE_METATYPE(settingsController15bc46*)


void settingsController15bc46_settingsController15bc46_QRegisterMetaTypes() {
}

void settingsController15bc46_ConnectUpdateSettings(void* ptr, long long t)
{
	QObject::connect(static_cast<settingsController15bc46*>(ptr), static_cast<void (settingsController15bc46::*)(QString, QString, QString)>(&settingsController15bc46::updateSettings), static_cast<settingsController15bc46*>(ptr), static_cast<void (settingsController15bc46::*)(QString, QString, QString)>(&settingsController15bc46::Signal_UpdateSettings), static_cast<Qt::ConnectionType>(t));
}

void settingsController15bc46_DisconnectUpdateSettings(void* ptr)
{
	QObject::disconnect(static_cast<settingsController15bc46*>(ptr), static_cast<void (settingsController15bc46::*)(QString, QString, QString)>(&settingsController15bc46::updateSettings), static_cast<settingsController15bc46*>(ptr), static_cast<void (settingsController15bc46::*)(QString, QString, QString)>(&settingsController15bc46::Signal_UpdateSettings));
}

void settingsController15bc46_UpdateSettings(void* ptr, struct Moc_PackedString id, struct Moc_PackedString name, struct Moc_PackedString description)
{
	static_cast<settingsController15bc46*>(ptr)->updateSettings(QString::fromUtf8(id.data, id.len), QString::fromUtf8(name.data, name.len), QString::fromUtf8(description.data, description.len));
}

int settingsController15bc46_settingsController15bc46_QRegisterMetaType()
{
	return qRegisterMetaType<settingsController15bc46*>();
}

int settingsController15bc46_settingsController15bc46_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<settingsController15bc46*>(const_cast<const char*>(typeName));
}

int settingsController15bc46_settingsController15bc46_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<settingsController15bc46>();
#else
	return 0;
#endif
}

int settingsController15bc46_settingsController15bc46_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<settingsController15bc46>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* settingsController15bc46___childItems_atList(void* ptr, int i)
{
	return ({QQuickItem * tmp = static_cast<QList<QQuickItem *>*>(ptr)->at(i); if (i == static_cast<QList<QQuickItem *>*>(ptr)->size()-1) { static_cast<QList<QQuickItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void settingsController15bc46___childItems_setList(void* ptr, void* i)
{
	static_cast<QList<QQuickItem *>*>(ptr)->append(static_cast<QQuickItem*>(i));
}

void* settingsController15bc46___childItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQuickItem *>();
}

int settingsController15bc46___grabTouchPoints_ids_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void settingsController15bc46___grabTouchPoints_ids_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* settingsController15bc46___grabTouchPoints_ids_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* settingsController15bc46___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void settingsController15bc46___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* settingsController15bc46___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* settingsController15bc46___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void settingsController15bc46___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* settingsController15bc46___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* settingsController15bc46___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void settingsController15bc46___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* settingsController15bc46___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* settingsController15bc46___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void settingsController15bc46___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* settingsController15bc46___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* settingsController15bc46___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void settingsController15bc46___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* settingsController15bc46___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* settingsController15bc46_NewSettingsController(void* parent)
{
		return new settingsController15bc46(static_cast<QQuickItem*>(parent));
}

void settingsController15bc46_DestroySettingsController(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->~settingsController15bc46();
}

void settingsController15bc46_DestroySettingsControllerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char settingsController15bc46_ChildMouseEventFilterDefault(void* ptr, void* item, void* event)
{
	return static_cast<settingsController15bc46*>(ptr)->QQuickItem::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

void settingsController15bc46_ClassBeginDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::classBegin();
}

void settingsController15bc46_ComponentCompleteDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::componentComplete();
}

char settingsController15bc46_ContainsDefault(void* ptr, void* point)
{
	return static_cast<settingsController15bc46*>(ptr)->QQuickItem::contains(*static_cast<QPointF*>(point));
}

void settingsController15bc46_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void settingsController15bc46_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void settingsController15bc46_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void settingsController15bc46_DropEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::dropEvent(static_cast<QDropEvent*>(event));
}

char settingsController15bc46_EventDefault(void* ptr, void* ev)
{
	return static_cast<settingsController15bc46*>(ptr)->QQuickItem::event(static_cast<QEvent*>(ev));
}

void settingsController15bc46_FocusInEventDefault(void* ptr, void* vqf)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::focusInEvent(static_cast<QFocusEvent*>(vqf));
}

void settingsController15bc46_FocusOutEventDefault(void* ptr, void* vqf)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::focusOutEvent(static_cast<QFocusEvent*>(vqf));
}

void settingsController15bc46_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void settingsController15bc46_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void settingsController15bc46_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void settingsController15bc46_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void settingsController15bc46_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* settingsController15bc46_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<settingsController15bc46*>(ptr)->QQuickItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char settingsController15bc46_IsTextureProviderDefault(void* ptr)
{
	return static_cast<settingsController15bc46*>(ptr)->QQuickItem::isTextureProvider();
}

void settingsController15bc46_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void settingsController15bc46_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void settingsController15bc46_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void settingsController15bc46_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void settingsController15bc46_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void settingsController15bc46_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void settingsController15bc46_MouseUngrabEventDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::mouseUngrabEvent();
}

void settingsController15bc46_ReleaseResourcesDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::releaseResources();
}

void* settingsController15bc46_TextureProviderDefault(void* ptr)
{
	return static_cast<settingsController15bc46*>(ptr)->QQuickItem::textureProvider();
}

void settingsController15bc46_TouchEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::touchEvent(static_cast<QTouchEvent*>(event));
}

void settingsController15bc46_TouchUngrabEventDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::touchUngrabEvent();
}

void settingsController15bc46_UpdateDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::update();
}

void settingsController15bc46_UpdatePolishDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::updatePolish();
}

void settingsController15bc46_WheelEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::wheelEvent(static_cast<QWheelEvent*>(event));
}

void settingsController15bc46_ChildEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::childEvent(static_cast<QChildEvent*>(event));
}

void settingsController15bc46_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void settingsController15bc46_CustomEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::customEvent(static_cast<QEvent*>(event));
}

void settingsController15bc46_DeleteLaterDefault(void* ptr)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::deleteLater();
}

void settingsController15bc46_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char settingsController15bc46_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<settingsController15bc46*>(ptr)->QQuickItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void settingsController15bc46_TimerEventDefault(void* ptr, void* event)
{
	static_cast<settingsController15bc46*>(ptr)->QQuickItem::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
