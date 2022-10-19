package flags

/*
	Line cap is style of line ending. i.e. It may be squared or rounded
*/

type LineCap byte

const (
	B_Butt   LineCap = 0x0
	B_Round  LineCap = 0x1
	B_Square LineCap = 0x2
)

/*
	Line join allows us to set type of place where lines connected
*/

type LineJoin byte

const (
	B_JoinMiter LineJoin = 0x0
	B_JoinRound LineJoin = 0x1
	B_JoinBevel LineJoin = 0x2
)

/*
	This sets the type of compositing operation to apply when drawing
	new shapes, where type is a string identifying which of the twelve
	compositing operations to use. Represented in bytes
*/

type CompositeOperation byte

const (
	B_SourceOver      CompositeOperation = 0x0
	B_SourceIn        CompositeOperation = 0x1
	B_SourceOut       CompositeOperation = 0x2
	B_SourceAtop      CompositeOperation = 0x3
	B_DestinationOver CompositeOperation = 0x4
	B_DestinationIn   CompositeOperation = 0x5
	B_DestinationOut  CompositeOperation = 0x6
	B_DestinationAtop CompositeOperation = 0x7
	B_Lighter         CompositeOperation = 0x8
	B_Copy            CompositeOperation = 0x9
	B_XOR             CompositeOperation = 0xA
	B_Multiply        CompositeOperation = 0xB
	B_Screen          CompositeOperation = 0xC
	B_Overlay         CompositeOperation = 0xD
	B_Darken          CompositeOperation = 0xE
	B_Lighten         CompositeOperation = 0xF
	B_ColorDodge      CompositeOperation = 0x10
	B_ColorBurn       CompositeOperation = 0x11
	B_HardLight       CompositeOperation = 0x12
	B_SoftLight       CompositeOperation = 0x13
	B_Difference      CompositeOperation = 0x14
	B_Exclusion       CompositeOperation = 0x15
	B_Hue             CompositeOperation = 0x16
	B_Saturation      CompositeOperation = 0x17
	B_Color           CompositeOperation = 0x18
	B_Luminosity      CompositeOperation = 0x19
)

/*
	Text Align show us how to positionate text in the field
*/

type TextAlign byte

const (
	B_AlignStart  TextAlign = 0x0
	B_AlignEnd    TextAlign = 0x1
	B_AlignLeft   TextAlign = 0x2
	B_AlignRight  TextAlign = 0x3
	B_AlignCenter TextAlign = 0x4
)

/*
	Text base line can be situated up, down, middle etc.
*/

type TextBaseline byte

const (
	B_lineAlphabetic  TextBaseline = 0x0
	B_lineIdeographic TextBaseline = 0x1
	B_lineTop         TextBaseline = 0x2
	B_lineBottom      TextBaseline = 0x3
	B_lineHanging     TextBaseline = 0x4
	B_lineMiddle      TextBaseline = 0x5
)

/*
	Pattern sets the something to be repeated
*/

type PatternRepetition byte

const (
	B_PatternRepeat   PatternRepetition = 0x0
	B_PatternRepeatX  PatternRepetition = 0x1
	B_PatternRepeatY  PatternRepetition = 0x2
	B_PatternNoRepeat PatternRepetition = 0x3
)

/*
	General draw commands of canvas
*/

type DrawCommand = byte

const (
	B_Arc                      DrawCommand = 0x0
	B_ArcTo                    DrawCommand = 0x1
	B_BeginPath                DrawCommand = 0x2
	B_BezierCurveTo            DrawCommand = 0x3
	B_ClearRect                DrawCommand = 0x4
	B_Clip                     DrawCommand = 0x5
	B_ClosePath                DrawCommand = 0x6
	B_CreateImageData          DrawCommand = 0x7
	B_CreateLinearGradient     DrawCommand = 0x8
	B_CreatePattern            DrawCommand = 0x9
	B_CreateRadialGradient     DrawCommand = 0xA
	B_DrawImage                DrawCommand = 0xB
	B_Ellipse                  DrawCommand = 0xC
	B_Fill                     DrawCommand = 0xD
	B_FillRect                 DrawCommand = 0xE
	B_FillStyle                DrawCommand = 0xF
	B_FillText                 DrawCommand = 0x10 // 16
	B_Font                     DrawCommand = 0x11
	B_GradientAddColorStop     DrawCommand = 0x12
	B_FillStyleGradient        DrawCommand = 0x13
	B_GlobalAlpha              DrawCommand = 0x14 // 20
	B_GlobalCompositeOperation DrawCommand = 0x15
	B_ImageSmoothingEnabled    DrawCommand = 0x16
	B_StrokeStyleGradient      DrawCommand = 0x17
	B_ReleasePattern           DrawCommand = 0x18
	B_LineCap                  DrawCommand = 0x19
	B_LineDashOffset           DrawCommand = 0x1A
	B_LineJoin                 DrawCommand = 0x1B
	B_LineTo                   DrawCommand = 0x1C
	B_LineWidth                DrawCommand = 0x1D
	B_ReleaseGradient          DrawCommand = 0x1E // 30
	B_MiterLimit               DrawCommand = 0x1F
	B_MoveTo                   DrawCommand = 0x20
	B_PutImageData             DrawCommand = 0x21
	B_QuadraticCurveTo         DrawCommand = 0x22
	B_Rect                     DrawCommand = 0x23
	B_Restore                  DrawCommand = 0x24
	B_Rotate                   DrawCommand = 0x25
	B_Save                     DrawCommand = 0x26
	B_Scale                    DrawCommand = 0x27
	B_PutLineDash              DrawCommand = 0x28 // 40
	B_PutTransform             DrawCommand = 0x29
	B_ShadowBlur               DrawCommand = 0x2A
	B_ShadowColor              DrawCommand = 0x2B
	B_ShadowOffsetX            DrawCommand = 0x2C
	B_ShadowOffsetY            DrawCommand = 0x2D
	B_Stroke                   DrawCommand = 0x2E
	B_StrokeRect               DrawCommand = 0x2F
	B_StrokeStyle              DrawCommand = 0x30
	B_StrokeText               DrawCommand = 0x31
	B_TextAlign                DrawCommand = 0x32 // 50
	B_TextBaseline             DrawCommand = 0x33
	B_Transform                DrawCommand = 0x34
	B_Translate                DrawCommand = 0x35
	B_FillTextMaxWidth         DrawCommand = 0x36
	B_StrokeTextMaxWidth       DrawCommand = 0x37
	B_PutImageDataDirty        DrawCommand = 0x38
	B_DrawImageScaled          DrawCommand = 0x39
	B_DrawImageSubRectangle    DrawCommand = 0x3A
	B_ReleaseImageData         DrawCommand = 0x3B
	B_FillStylePattern         DrawCommand = 0x3C
	B_StrokeStylePattern       DrawCommand = 0x3D
	B_GetImageData             DrawCommand = 0x3E // 62
)
