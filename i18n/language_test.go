package i18n

import (
	"fmt"
)

func ExampleAllLanguages() {
	for i, lang := range AllLanguages() {
		fmt.Println(i, lang.Code, lang.NativeName)
	}
	//Output:
	//0 aa Afaraf
	//1 ab Аҧсуа
	//2 ae avesta
	//3 af Afrikaans
	//4 ak Akan
	//5 am አማርኛ
	//6 an Aragonés
	//7 ar العربية
	//8 as অসমীয়া
	//9 av авар мацӀ
	//10 ay aymar aru
	//11 az azərbaycan dili
	//12 ba башҡорт теле
	//13 be Беларуская
	//14 bg български език
	//15 bh भोजपुरी
	//16 bi Bislama
	//17 bm bamanankan
	//18 bn বাংলা
	//19 bo བོད་ཡིག
	//20 br brezhoneg
	//21 bs bosanski jezik
	//22 ca Català
	//23 ce нохчийн мотт
	//24 ch Chamoru
	//25 co corsu
	//26 cr ᓀᐦᐃᔭᐍᐏᐣ
	//27 cs česky
	//28 cu ѩзыкъ словѣньскъ
	//29 cv чӑваш чӗлхи
	//30 cy Cymraeg
	//31 da dansk
	//32 de Deutsch
	//33 dv ދިވެހި
	//34 dz རྫོང་ཁ
	//35 ee Ɛʋɛgbɛ
	//36 el Ελληνικά
	//37 en English
	//38 eo Esperanto
	//39 es español
	//40 et Eesti keel
	//41 eu euskara
	//42 fa فارسی
	//43 ff Fulfulde
	//44 fi suomen kieli
	//45 fj vosa Vakaviti
	//46 fo Føroyskt
	//47 fr français
	//48 fy Frysk
	//49 ga Gaeilge
	//50 gd Gàidhlig
	//51 gl Galego
	//52 gn Avañe'ẽ
	//53 gu ગુજરાતી
	//54 gv Ghaelg
	//55 ha هَوُسَ
	//56 he עברית
	//57 hi हिन्दी
	//58 ho Hiri Motu
	//59 hr Hrvatski
	//60 ht Kreyòl ayisyen
	//61 hu Magyar
	//62 hy Հայերեն
	//63 hz Otjiherero
	//64 ia Interlingua
	//65 id Bahasa Indonesia
	//66 ie Interlingue
	//67 ig Igbo
	//68 ii ꆇꉙ
	//69 ik Iñupiaq
	//70 io Ido
	//71 is Íslenska
	//72 it Italiano
	//73 iu ᐃᓄᒃᑎᑐᑦ
	//74 ja 日本語
	//75 jv basa Jawa
	//76 ka ქართული
	//77 kg KiKongo
	//78 ki Gĩkũyũ
	//79 kj Kuanyama
	//80 kk Қазақ тілі
	//81 kl kalaallisut
	//82 km ភាសាខ្មែរ
	//83 kn ಕನ್ನಡ
	//84 ko 한국어
	//85 kr Kanuri
	//86 ks कश्मीरी
	//87 ku Kurdî
	//88 kv коми кыв
	//89 kw Kernewek
	//90 ky кыргыз тили
	//91 la latine
	//92 lb Lëtzebuergesch
	//93 lg Luganda
	//94 li Limburgs
	//95 ln Lingála
	//96 lo ພາສາລາວ
	//97 lt lietuvių kalba
	//98 lu Tshiluba
	//99 lv latviešu valoda
	//100 mg Malagasy fiteny
	//101 mh Kajin M̧ajeļ
	//102 mi te reo Māori
	//103 mk македонски јазик
	//104 ml മലയാളം
	//105 mn ᠮᠣᠩᠭᠣᠯ ᠬᠡᠯᠡ
	//106 mo лимба молдовеняскэ
	//107 mr मराठी
	//108 ms bahasa Melayu
	//109 mt Malti
	//110 my ဗမာစာ
	//111 na Ekakairũ Naoero
	//112 nb Norsk bokmål
	//113 nd isiNdebele
	//114 ne नेपाली
	//115 ng Owambo
	//116 nl Nederlands
	//117 nn Norsk nynorsk
	//118 no Norsk
	//119 nr Ndébélé
	//120 nv Diné bizaad
	//121 ny chiCheŵa
	//122 oc Occitan
	//123 oj ᐊᓂᔑᓈᐯᒧᐎᓐ
	//124 om Afaan Oromoo
	//125 or ଓଡ଼ିଆ
	//126 os Ирон æвзаг
	//127 pa ਪੰਜਾਬੀ
	//128 pi पाऴि
	//129 pl polski
	//130 ps پښتو
	//131 pt Português
	//132 qu Runa Simi
	//133 rm rumantsch grischun
	//134 rn kiRundi
	//135 ro română
	//136 ru русский язык
	//137 rw Kinyarwanda
	//138 sa संस्कृतम्
	//139 sc sardu
	//140 sd सिन्धी
	//141 se Davvisámegiella
	//142 sg yângâ tî sängö
	//143 sh Srpskohrvatski
	//144 si සිංහල
	//145 sk slovenčina
	//146 sl slovenščina
	//147 sm gagana fa'a Samoa
	//148 sn chiShona
	//149 so Soomaaliga
	//150 sq Shqip
	//151 sr српски језик
	//152 ss SiSwati
	//153 st seSotho
	//154 su Basa Sunda
	//155 sv Svenska
	//156 sw Kiswahili
	//157 ta தமிழ்
	//158 te తెలుగు
	//159 tg тоҷикӣ
	//160 th ไทย
	//161 ti ትግርኛ
	//162 tk Türkmen
	//163 tl Tagalog
	//164 tn seTswana
	//165 to faka Tonga
	//166 tr Türkçe
	//167 ts xiTsonga
	//168 tt татарча
	//169 tw Twi
	//170 ty Reo Mā`ohi
	//171 ug ئۇيغۇرچە
	//172 uk Українська
	//173 ur اردو
	//174 uz O'zbek
	//175 ve tshiVenḓa
	//176 vi Tiếng Việt
	//177 vo Volapük
	//178 wa Walon
	//179 wo Wollof
	//180 xh isiXhosa
	//181 yi ייִדיש
	//182 yo Yorùbá
	//183 za Saw cuengh
	//184 zh 中文
	//185 zh-hans 简体中文
	//186 zh-hant 繁體中文
	//187 zu isiZulu
}
