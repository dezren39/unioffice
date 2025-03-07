//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package terms ;import (_g "encoding/xml";_c "fmt";_a "github.com/unidoc/unioffice";_cf "github.com/unidoc/unioffice/common/logger";_af "github.com/unidoc/unioffice/schema/purl.org/dc/elements";);func (_abe *MESH )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_bgg ,_bcc :=d .Token ();
if _bcc !=nil {return _c .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073",_bcc );};if _adf ,_ccb :=_bgg .(_g .EndElement );_ccb &&_adf .Name ==start .Name {break ;};};return nil ;};func (_eee *ElementOrRefinementContainer )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072";
e .EncodeToken (start );if _eee .Choice !=nil {for _ ,_ec :=range _eee .Choice {_ec .MarshalXML (e ,_g .StartElement {});};};e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func NewBox ()*Box {_e :=&Box {};return _e };

// Validate validates the ISO3166 and its children
func (_cfe *ISO3166 )Validate ()error {return _cfe .ValidateWithPath ("\u0049S\u004f\u0033\u0031\u0036\u0036");};

// Validate validates the DCMIType and its children
func (_gd *DCMIType )Validate ()error {return _gd .ValidateWithPath ("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065");};func (_fbe *LCSH )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_gaa ,_affg :=d .Token ();if _affg !=nil {return _c .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073",_affg );
};if _ce ,_ebc :=_gaa .(_g .EndElement );_ebc &&_ce .Name ==start .Name {break ;};};return nil ;};func (_acb *Period )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_bbd ,_bda :=d .Token ();if _bda !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073",_bda );
};if _cff ,_efg :=_bbd .(_g .EndElement );_efg &&_cff .Name ==start .Name {break ;};};return nil ;};

// Validate validates the ElementsAndRefinementsGroup and its children
func (_bge *ElementsAndRefinementsGroup )Validate ()error {return _bge .ValidateWithPath ("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070");};

// Validate validates the LCSH and its children
func (_aaf *LCSH )Validate ()error {return _aaf .ValidateWithPath ("\u004c\u0043\u0053\u0048")};func NewDCMIType ()*DCMIType {_gg :=&DCMIType {};return _gg };func (_dae *ElementsAndRefinementsGroup )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _dae .Choice !=nil {for _ ,_ccc :=range _dae .Choice {_ccc .MarshalXML (e ,_g .StartElement {});
};};return nil ;};

// Validate validates the Period and its children
func (_ebb *Period )Validate ()error {return _ebb .ValidateWithPath ("\u0050\u0065\u0072\u0069\u006f\u0064");};

// Validate validates the RFC1766 and its children
func (_bfd *RFC1766 )Validate ()error {return _bfd .ValidateWithPath ("\u0052F\u0043\u0031\u0037\u0036\u0036");};type TGN struct{};func (_fce *ISO3166 )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0049S\u004f\u0033\u0031\u0036\u0036";
e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func NewRFC3066 ()*RFC3066 {_eeg :=&RFC3066 {};return _eeg };func (_ff *TGN )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0054\u0047\u004e";
e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func NewIMT ()*IMT {_bac :=&IMT {};return _bac };func (_aag *Point )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_efa ,_ddb :=d .Token ();if _ddb !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073",_ddb );
};if _fgb ,_dbe :=_efa .(_g .EndElement );_dbe &&_fgb .Name ==start .Name {break ;};};return nil ;};func (_fba *ElementsAndRefinementsGroupChoice )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_dge :for {_ggfb ,_fbb :=d .Token ();if _fbb !=nil {return _fbb ;
};switch _bff :=_ggfb .(type ){case _g .StartElement :switch _bff .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cb :=_af .NewAny ();
if _fa :=d .DecodeElement (_cb ,&_bff );_fa !=nil {return _fa ;};_fba .Any =append (_fba .Any ,_cb );default:_cf .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076",_bff .Name );
if _fc :=d .Skip ();_fc !=nil {return _fc ;};};case _g .EndElement :break _dge ;case _g .CharData :};};return nil ;};type ISO639_2 struct{};type ISO3166 struct{};func (_dcg *LCSH )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u004c\u0043\u0053\u0048";
e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func NewDDC ()*DDC {_bd :=&DDC {};return _bd };func NewMESH ()*MESH {_ac :=&MESH {};return _ac };

// Validate validates the ISO639_2 and its children
func (_baa *ISO639_2 )Validate ()error {return _baa .ValidateWithPath ("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032");};func (_bbg *ElementsAndRefinementsGroupChoice )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _bbg .Any !=nil {_bgf :=_g .StartElement {Name :_g .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_gdf :=range _bbg .Any {e .EncodeElement (_gdf ,_bgf );};};return nil ;};type IMT struct{};

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_cffd *RFC3066 )ValidateWithPath (path string )error {return nil };func NewPeriod ()*Period {_aabe :=&Period {};return _aabe };func (_be *DCMIType )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_aab ,_fe :=d .Token ();if _fe !=nil {return _c .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073",_fe );
};if _feg ,_cfa :=_aab .(_g .EndElement );_cfa &&_feg .Name ==start .Name {break ;};};return nil ;};type UDC struct{};func (_cadc *ISO639_2 )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032";
e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};type ElementsAndRefinementsGroup struct{Choice []*ElementsAndRefinementsGroupChoice ;};func (_ea *DCMIType )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065";
e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func (_eff *IMT )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_dcc ,_fd :=d .Token ();if _fd !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073",_fd );
};if _bcd ,_eaf :=_dcc .(_g .EndElement );_eaf &&_bcd .Name ==start .Name {break ;};};return nil ;};func (_b *Box )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_cc ,_aa :=d .Token ();if _aa !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073",_aa );
};if _f ,_ba :=_cc .(_g .EndElement );_ba &&_f .Name ==start .Name {break ;};};return nil ;};func (_dd *ISO3166 )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_bbe ,_bcf :=d .Token ();if _bcf !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073",_bcf );
};if _cae ,_dea :=_bbe .(_g .EndElement );_dea &&_cae .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_afd *ISO3166 )ValidateWithPath (path string )error {return nil };func (_ca *Box )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0042\u006f\u0078";e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });
return nil ;};func (_agg *ISO639_2 )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_bag ,_cccd :=d .Token ();if _cccd !=nil {return _c .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073",_cccd );
};if _bga ,_gb :=_bag .(_g .EndElement );_gb &&_bga .Name ==start .Name {break ;};};return nil ;};func (_cfg *IMT )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0049\u004d\u0054";e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });
return nil ;};

// Validate validates the UDC and its children
func (_egb *UDC )Validate ()error {return _egb .ValidateWithPath ("\u0055\u0044\u0043")};

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_cad *ElementsAndRefinementsGroupChoice )ValidateWithPath (path string )error {for _agf ,_cce :=range _cad .Any {if _gac :=_cce .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_agf ));_gac !=nil {return _gac ;
};};return nil ;};

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_ga *ElementsAndRefinementsGroup )ValidateWithPath (path string )error {for _fge ,_eae :=range _ga .Choice {if _dac :=_eae .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_fge ));
_dac !=nil {return _dac ;};};return nil ;};func (_ee *DDC )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_ccg ,_bb :=d .Token ();if _bb !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073",_bb );
};if _ef ,_bea :=_ccg .(_g .EndElement );_bea &&_ef .Name ==start .Name {break ;};};return nil ;};

// Validate validates the IMT and its children
func (_dgf *IMT )Validate ()error {return _dgf .ValidateWithPath ("\u0049\u004d\u0054")};func (_bacg *RFC1766 )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_dabb ,_adb :=d .Token ();if _adb !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073",_adb );
};if _cca ,_caeb :=_dabb .(_g .EndElement );_caeb &&_cca .Name ==start .Name {break ;};};return nil ;};func (_abd *RFC3066 )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_eef ,_geb :=d .Token ();if _geb !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073",_geb );
};if _ecc ,_fdf :=_eef .(_g .EndElement );_fdf &&_ecc .Name ==start .Name {break ;};};return nil ;};func NewURI ()*URI {_befc :=&URI {};return _befc };type Period struct{};

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_cd *W3CDTF )ValidateWithPath (path string )error {return nil };type Point struct{};

// Validate validates the Point and its children
func (_bfa *Point )Validate ()error {return _bfa .ValidateWithPath ("\u0050\u006f\u0069n\u0074")};

// Validate validates the URI and its children
func (_bee *URI )Validate ()error {return _bee .ValidateWithPath ("\u0055\u0052\u0049")};type W3CDTF struct{};

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_gf *DCMIType )ValidateWithPath (path string )error {return nil };type RFC1766 struct{};func NewElementsAndRefinementsGroup ()*ElementsAndRefinementsGroup {_fee :=&ElementsAndRefinementsGroup {};return _fee ;};type ElementsAndRefinementsGroupChoice struct{Any []*_af .Any ;
};func (_afg *URI )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0055\u0052\u0049";e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_fbbe *LCC )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_ead *TGN )ValidateWithPath (path string )error {return nil };func (_ggff *UDC )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_aca ,_bgfa :=d .Token ();if _bgfa !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073",_bgfa );
};if _gbd ,_bead :=_aca .(_g .EndElement );_bead &&_gbd .Name ==start .Name {break ;};};return nil ;};func (_deb *Period )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0050\u0065\u0072\u0069\u006f\u0064";e .EncodeToken (start );
e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func (_fbae *RFC1766 )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0052F\u0043\u0031\u0037\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });
return nil ;};

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_fea *IMT )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_bbgb *ISO639_2 )ValidateWithPath (path string )error {return nil };func NewPoint ()*Point {_gdfg :=&Point {};return _gdfg };func (_daad *URI )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_dbf ,_fdee :=d .Token ();if _fdee !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073",_fdee );
};if _adba ,_gba :=_dbf .(_g .EndElement );_gba &&_adba .Name ==start .Name {break ;};};return nil ;};

// Validate validates the LCC and its children
func (_ad *LCC )Validate ()error {return _ad .ValidateWithPath ("\u004c\u0043\u0043")};func NewW3CDTF ()*W3CDTF {_ddg :=&W3CDTF {};return _ddg };func NewTGN ()*TGN {_bbc :=&TGN {};return _bbc };

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_ab *Box )ValidateWithPath (path string )error {return nil };func (_daf *LCC )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u004c\u0043\u0043";e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });
return nil ;};func NewISO3166 ()*ISO3166 {_df :=&ISO3166 {};return _df };func (_cg *UDC )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0055\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });
return nil ;};

// Validate validates the DDC and its children
func (_bdc *DDC )Validate ()error {return _bdc .ValidateWithPath ("\u0044\u0044\u0043")};type DDC struct{};func (_fb *ElementOrRefinementContainer )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_bdd :for {_dc ,_afe :=d .Token ();if _afe !=nil {return _afe ;
};switch _de :=_dc .(type ){case _g .StartElement :switch _de .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_gfd :=NewElementsAndRefinementsGroupChoice ();
if _abc :=d .DecodeElement (&_gfd .Any ,&_de );_abc !=nil {return _abc ;};_fb .Choice =append (_fb .Choice ,_gfd );default:_cf .Log .Debug ("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076",_de .Name );
if _dcb :=d .Skip ();_dcb !=nil {return _dcb ;};};case _g .EndElement :break _bdd ;case _g .CharData :};};return nil ;};type URI struct{};type LCSH struct{};type ElementOrRefinementContainer struct{Choice []*ElementsAndRefinementsGroupChoice ;};func (_fga *MESH )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u004d\u0045\u0053\u0048";
e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func NewRFC1766 ()*RFC1766 {_daaa :=&RFC1766 {};return _daaa };

// Validate validates the TGN and its children
func (_bdfd *TGN )Validate ()error {return _bdfd .ValidateWithPath ("\u0054\u0047\u004e")};

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_abb *UDC )ValidateWithPath (path string )error {return nil };type RFC3066 struct{};func NewElementOrRefinementContainer ()*ElementOrRefinementContainer {_aba :=&ElementOrRefinementContainer {};return _aba ;};

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_daa *Period )ValidateWithPath (path string )error {return nil };func NewLCSH ()*LCSH {_deg :=&LCSH {};return _deg };

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_eb *ElementOrRefinementContainer )ValidateWithPath (path string )error {for _da ,_bc :=range _eb .Choice {if _bdca :=_bc .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_da ));
_bdca !=nil {return _bdca ;};};return nil ;};type Box struct{};

// Validate validates the MESH and its children
func (_beff *MESH )Validate ()error {return _beff .ValidateWithPath ("\u004d\u0045\u0053\u0048")};func (_aac *W3CDTF )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0057\u0033\u0043\u0044\u0054\u0046";e .EncodeToken (start );
e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_bef *LCSH )ValidateWithPath (path string )error {return nil };func NewElementsAndRefinementsGroupChoice ()*ElementsAndRefinementsGroupChoice {_dgb :=&ElementsAndRefinementsGroupChoice {};return _dgb ;};type LCC struct{};func (_cab *TGN )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_age ,_dfc :=d .Token ();
if _dfc !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073",_dfc );};if _ggb ,_befe :=_age .(_g .EndElement );_befe &&_ggb .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_ag *DDC )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_fdc *MESH )ValidateWithPath (path string )error {return nil };

// Validate validates the ElementOrRefinementContainer and its children
func (_aff *ElementOrRefinementContainer )Validate ()error {return _aff .ValidateWithPath ("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072");};func (_ggf *DDC )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0044\u0044\u0043";
e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};

// Validate validates the Box and its children
func (_bg *Box )Validate ()error {return _bg .ValidateWithPath ("\u0042\u006f\u0078")};

// Validate validates the W3CDTF and its children
func (_caf *W3CDTF )Validate ()error {return _caf .ValidateWithPath ("\u0057\u0033\u0043\u0044\u0054\u0046");};func (_acd *RFC3066 )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0052F\u0043\u0033\u0030\u0036\u0036";e .EncodeToken (start );
e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func (_gef *Point )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0050\u006f\u0069n\u0074";e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });
return nil ;};func NewUDC ()*UDC {_bgab :=&UDC {};return _bgab };

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_bdf *RFC1766 )ValidateWithPath (path string )error {return nil };func (_ebg *ElementsAndRefinementsGroup )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_fg :for {_ccga ,_eg :=d .Token ();if _eg !=nil {return _eg ;};switch _gga :=_ccga .(type ){case _g .StartElement :switch _gga .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_ebge :=NewElementsAndRefinementsGroupChoice ();
if _fegd :=d .DecodeElement (&_ebge .Any ,&_gga );_fegd !=nil {return _fegd ;};_ebg .Choice =append (_ebg .Choice ,_ebge );default:_cf .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076",_gga .Name );
if _gc :=d .Skip ();_gc !=nil {return _gc ;};};case _g .EndElement :break _fg ;case _g .CharData :};};return nil ;};

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_dcbg *URI )ValidateWithPath (path string )error {return nil };type DCMIType struct{};

// Validate validates the RFC3066 and its children
func (_dega *RFC3066 )Validate ()error {return _dega .ValidateWithPath ("\u0052F\u0043\u0033\u0030\u0036\u0036");};func (_gbe *W3CDTF )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_cag ,_ggffc :=d .Token ();if _ggffc !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073",_ggffc );
};if _gae ,_cgb :=_cag .(_g .EndElement );_cgb &&_gae .Name ==start .Name {break ;};};return nil ;};func NewLCC ()*LCC {_gad :=&LCC {};return _gad };func (_feee *LCC )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_ge ,_fde :=d .Token ();
if _fde !=nil {return _c .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073",_fde );};if _eaa ,_bddf :=_ge .(_g .EndElement );_bddf &&_eaa .Name ==start .Name {break ;};};return nil ;};type MESH struct{};

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_cbd *ElementsAndRefinementsGroupChoice )Validate ()error {return _cbd .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065");
};

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_dab *Point )ValidateWithPath (path string )error {return nil };func NewISO639_2 ()*ISO639_2 {_afb :=&ISO639_2 {};return _afb };func init (){_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0053\u0048",NewLCSH );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004d\u0045\u0053\u0048",NewMESH );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0044\u0043",NewDDC );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0043",NewLCC );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0044\u0043",NewUDC );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u0065\u0072\u0069\u006f\u0064",NewPeriod );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0057\u0033\u0043\u0044\u0054\u0046",NewW3CDTF );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065",NewDCMIType );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u004d\u0054",NewIMT );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0052\u0049",NewURI );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032",NewISO639_2 );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0031\u0037\u0036\u0036",NewRFC1766 );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0033\u0030\u0036\u0036",NewRFC3066 );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u006f\u0069n\u0074",NewPoint );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049S\u004f\u0033\u0031\u0036\u0036",NewISO3166 );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0042\u006f\u0078",NewBox );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0054\u0047\u004e",NewTGN );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072",NewElementOrRefinementContainer );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070",NewElementsAndRefinementsGroup );
};