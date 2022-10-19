package flags

/*
	Line cap is style of line ending. i.e. It may be squared or rounded
*/

type LineCap byte

const (
	bButt LineCap = iota
	bRound
	bSquare
)

/*
	Line join allows us to set type of place where lines connected
*/

type LineJoin byte

const (
	bJoinMiter LineJoin = iota
	bJoinRound
	bJoinBevel
)

/*
	This sets the type of compositing operation to apply when drawing
	new shapes, where type is a string identifying which of the twelve
	compositing operations to use. Represented in bytes
*/

type CompositeOperation byte

const (
	bSourceOver CompositeOperation = iota
	bSourceIn
	bSourceOut
	bSourceAtop
	bDestinationOver
	bDestinationIn
	bDestinationOut
	bDestinationAtop
	bLighter
	bCopy
	bXOR
	bMultiply
	bScreen
	bOverlay
	bDarken
	bLighten
	bColorDodge
	bColorBurn
	bHardLight
	bSoftLight
	bDifference
	bExclusion
	bHue
	bSaturation
	bColor
	bLuminosity
)

/*
	Text Align show us how to positionate text in the field
*/

type TextAlign byte

const (
	bAlignStart TextAlign = iota
	bAlignEnd
	bAlignLeft
	bAlignRight
	bAlignCenter
)

/*
	Text base line can be situated up, down, middle etc.
*/

type TextBaseline byte

const (
	blineAlphabetic TextBaseline = iota
	blineIdeographic
	blineTop
	blineBottom
	blineHanging
	blineMiddle
)

/*
	Pattern sets the something to be repeated
*/

type PatternRepetition byte

const (
	bPatternRepeat PatternRepetition = iota
	bPatternRepeatX
	bPatternRepeatY
	bPatternNoRepeat
)

/*
	General draw commands of canvas
*/

const (
	bArc byte = 1 + iota
	bArcTo
	bBeginPath
	bBezierCurveTo
	bClearRect
	bClip
	bClosePath
	bCreateImageData
	bCreateLinearGradient
	bCreatePattern
	bCreateRadialGradient
	bDrawImage
	bEllipse
	bFill
	bFillRect
	bFillStyle
	bFillText
	bFont
	bGradientAddColorStop
	bGradientAddColorStopString
	bFillStyleGradient
	bGlobalAlpha
	bGlobalCompositeOperation
	bImageSmoothingEnabled
	bStrokeStyleGradient
	bReleasePattern
	bLineCap
	bLineDashOffset
	bLineJoin
	bLineTo
	bLineWidth
	bReleaseGradient
	bMiterLimit
	bMoveTo
	bPutImageData
	bQuadraticCurveTo
	bRect
	bRestore
	bRotate
	bSave
	bScale
	bSetLineDash
	bSetTransform
	bShadowBlur
	bShadowColor
	bShadowOffsetX
	bShadowOffsetY
	bStroke
	bStrokeRect
	bStrokeStyle
	bStrokeText
	bTextAlign
	bTextBaseline
	bTransform
	bTranslate
	bFillTextMaxWidth
	bStrokeTextMaxWidth
	bFillStyleString
	bStrokeStyleString
	bShadowColorString
	bPutImageDataDirty
	bDrawImageScaled
	bDrawImageSubRectangle
	bReleaseImageData
	bFillStylePattern
	bStrokeStylePattern
	bGetImageData
)
