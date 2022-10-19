package flags

/*
	Line cap is style of line ending. i.e. It may be squared or rounded
*/

type LineCap byte

const (
	B_Butt LineCap = iota
	B_Round
	B_Square
)

/*
	Line join allows us to set type of place where lines connected
*/

type LineJoin byte

const (
	B_JoinMiter LineJoin = iota
	B_JoinRound
	B_JoinBevel
)

/*
	This sets the type of compositing operation to apply when drawing
	new shapes, where type is a string identifying which of the twelve
	compositing operations to use. Represented in bytes
*/

type CompositeOperation byte

const (
	B_SourceOver CompositeOperation = iota
	B_SourceIn
	B_SourceOut
	B_SourceAtop
	B_DestinationOver
	B_DestinationIn
	B_DestinationOut
	B_DestinationAtop
	B_Lighter
	B_Copy
	B_XOR
	B_Multiply
	B_Screen
	B_Overlay
	B_Darken
	B_Lighten
	B_ColorDodge
	B_ColorBurn
	B_HardLight
	B_SoftLight
	B_Difference
	B_Exclusion
	B_Hue
	B_Saturation
	B_Color
	B_Luminosity
)

/*
	Text Align show us how to positionate text in the field
*/

type TextAlign byte

const (
	B_AlignStart TextAlign = iota
	B_AlignEnd
	B_AlignLeft
	B_AlignRight
	B_AlignCenter
)

/*
	Text base line can be situated up, down, middle etc.
*/

type TextBaseline byte

const (
	B_lineAlphabetic TextBaseline = iota
	B_lineIdeographic
	B_lineTop
	B_lineBottom
	B_lineHanging
	B_lineMiddle
)

/*
	Pattern sets the something to be repeated
*/

type PatternRepetition byte

const (
	B_PatternRepeat PatternRepetition = iota
	B_PatternRepeatX
	B_PatternRepeatY
	B_PatternNoRepeat
)

/*
	General draw commands of canvas
*/

const (
	B_Arc byte = 1 + iota
	B_ArcTo
	B_BeginPath
	B_BezierCurveTo
	B_ClearRect
	B_Clip
	B_ClosePath
	B_CreateImageData
	B_CreateLinearGradient
	B_CreatePattern
	B_CreateRadialGradient
	B_DrawImage
	B_Ellipse
	B_Fill
	B_FillRect
	B_FillStyle
	B_FillText
	B_Font
	B_GradientAddColorStop
	B_GradientAddColorStopString
	B_FillStyleGradient
	B_GlobalAlpha
	B_GlobalCompositeOperation
	B_ImageSmoothingEnabled
	B_StrokeStyleGradient
	B_ReleasePattern
	B_LineCap
	B_LineDashOffset
	B_LineJoin
	B_LineTo
	B_LineWidth
	B_ReleaseGradient
	B_MiterLimit
	B_MoveTo
	B_PutImageData
	B_QuadraticCurveTo
	B_Rect
	B_Restore
	B_Rotate
	B_Save
	B_Scale
	B_SetLineDash
	B_SetTransform
	B_ShadowBlur
	B_ShadowColor
	B_ShadowOffsetX
	B_ShadowOffsetY
	B_Stroke
	B_StrokeRect
	B_StrokeStyle
	B_StrokeText
	B_TextAlign
	B_TextBaseline
	B_Transform
	B_Translate
	B_FillTextMaxWidth
	B_StrokeTextMaxWidth
	B_FillStyleString
	B_StrokeStyleString
	B_ShadowColorString
	B_PutImageDataDirty
	B_DrawImageScaled
	B_DrawImageSubRectangle
	B_ReleaseImageData
	B_FillStylePattern
	B_StrokeStylePattern
	B_GetImageData
)
