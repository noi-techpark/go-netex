package netex

const XsiNamespace = "http://www.w3.org/2001/XMLSchema-instance"
const NetexNamespace = "http://www.netex.org.uk/netex"

// See section 4.2.2 of Italian netex profile

// Epip types

const EpipTypeCommon = "EU_PI_COMMON"
const EpipTypeMetadata = "EU_PI_METADATA"
const EpipTypeStop = "EU_PI_STOP"
const EpipTypeNetwork = "EU_PI_NETWORK"
const EpipTypeTimetable = "EU_PI_TIMETABLE"
const EpipTypeCalendar = "EU_PI_CALENDAR"
const EpipTypeLineOffer = "EU_PI_LINE_OFFER"
const EpipTypeNetworkOffer = "EU_PI_NETWORK_OFFER"
const EpipTypeStopOffer = "EU_PI_STOP_OFFER"
const EpipTypeMobility = "EU_PI_MOBILITY"

// Specific frame + valid epip type

const TypeResourceFrameCommon = "ResourceFrame-" + EpipTypeCommon
const TypeResourceFrameMetadata = "ResourceFrame-" + EpipTypeMetadata
const TypeSiteFrameStop = "SiteFrame-" + EpipTypeStop
const TypeServiceFrameNetwork = "ServiceFrame-" + EpipTypeNetwork
const TypeTimetableFrameTimetable = "TimetableFrame-" + EpipTypeTimetable
const TypeServiceCalendarFrameCalendar = "ServiceCalendarFrame-" + EpipTypeCalendar
const TypeMobilityServiceFrameMobility = "MobilityServiceFrame-" + EpipTypeMobility

// Composite frame + valid epip type

const TypeCompositeFrameLineOffer = "CompositeFrame-" + EpipTypeLineOffer
const TypeCompositeFrameNetworkOffer = "CompositeFrame-" + EpipTypeNetworkOffer
const TypeCompositeFrameStopOffer = "CompositeFrame-" + EpipTypeStopOffer
