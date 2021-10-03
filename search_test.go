package tvdb

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
)

func TestClient_Search(t *testing.T) {
	fireflyResult := `{
		"status": "success",
		"data": [
		  {
			"objectID": "series-78874",
			"aliases": [
			  "Firefly - Der Aufbruch der Serenity",
			  "Serenity"
			],
			"country": "usa",
			"id": "series-78874",
			"image_url": "https://artworks.thetvdb.com/banners/posters/78874-2.jpg",
			"name": "Firefly",
			"name_translated": "{\"ces\": \"Firefly\",\n\"zho\": \"萤火虫\",\n\"spa\": \"Firefly\",\n\"hun\": \"Firefly\",\n\"rus\": \"Светлячок\",\n\"heb\": \"פיירפלי\",\n\"pol\": \"Firefly\",\n\"swe\": \"Firefly\",\n\"hrv\": \"Firefly\",\n\"deu\": \"Firefly\",\n\"fra\": \"Firefly\",\n\"nld\": \"Firefly\",\n\"por\": \"Firefly\",\n\"ita\": \"Firefly\",\n\"kor\": \"파이어플라이\"\n}",
			"first_air_time": "2002-09-20",
			"overview": "In the far-distant future, Captain Malcolm \"Mal\" Reynolds is a renegade former brown-coat sergeant, now turned smuggler & rogue, who is the commander of a small spacecraft, with a loyal hand-picked crew made up of the first mate, Zoe Warren; the pilot Hoban \"Wash\" Washburn; the gung-ho grunt Jayne Cobb; the engineer Kaylee Frye; the fugitives Dr. Simon Tam and his psychic sister River. Together, they travel the far reaches of space in search of food, money, and anything to live on.",
			"overview_translated": [
			  "ces: Firefly je dechberoucím příběhem vesmírných psanců, kteří se pohybují na okraji společnosti, živí se převážně pašováním a jejich cílem není spasit galaxii, ale prostě přežít - \"letět dál\". Život jim komplikují jak vesmírné lodě totalitní Aliance, tak barbarští Plenitelé, sužující světy na okraji vesmíru. Hlavní síla seriálu spočívá v propracovaných charakterech postav, dobře vykonstruovaném příběhu a překvapivém rozuzlení jednotlivých epizod. Každý z hrdinů má své vlastní naděje a obavy a jejich osobnosti a vztahy se neustále vyvíjejí. ",
			  "deu: Im 26. Jahrhundert schlägt sich die Crew des Raumfrachters \"Serenity\" vorwiegend mit Schmuggelaufträgen durch den Weltraum. Aliens wurden bisher keine entdeckt, aber Captain Malcolm \"Mal\" Reynolds, ein ehemaliger Kämpfer im verlorenen Bürgerkrieg gegen die jetzt allgegenwärtige Allianz, hat trotzdem laufend Ärger - speziell nachdem er den jungen Arzt Simon und seine Schwester River an Bord nimmt. Jene ist das geflohene \"Produkt\" von Experimenten der Allianz, und wird von dieser überall gesucht.",
			  "fra: Après une guerre civile qui a permis à l'Union des Planètes de dominer l'espace, le capitaine Malcolm Reynolds et son équipage s'efforcent de survivre à bord du vaisseau Serenity en effectuant diverses missions (transport de marchandises ou une mission de sauvetage...) sans trop se soucier de la légalité.",
			  "heb: ג'וס ווידון, היוצר-תסריטאי-במאי-מפיק של \"באפי, ציידת הערפדים\" ו\"אנג'ל\", מפתיע מחדש עם סדרה שהיא הכלאה בלתי אפשרית בין מדע-בדיוני למערבון. בעתיד הקצת רחוק מסתבך הצוות של ספינת החלל \"סריניטי\" עם טרמפיסט מבוקש.",
			  "hun: A Firefly egy amerikai sci-fi sorozat, amely a műfajban megszokottnál kevesebb fantasztikus, futurisztikus elemet tartalmaz, inkább hétköznapi emberi és társadalmi problémákkal teli világot mutat be, egy másik műfaj, a western témáit is felhasználva. A történet szerint a 25-26. századra a Föld készletei már rég kimerültek, az emberiség a Naprendszeren kívüli világokban telepedett le. Az űr meghódításában Kína és az USA játszott vezető szerepet, amely két szuperhatalom alkotja a Szövetséget és az Egyesítési Háborúban közös kormányzat alatt egyesítette a belső bolygókat. A külső bolygókon, sok helyen a Vadnyugatra emlékeztető, jóval rosszabb körülmények között élnek az emberek.\r\n\r\nA sorozatnak - bár csak egy évadot élt meg - hatalmas rajongótábora van.",
			  "ita: Ambientata in un futuro privo di una reale libertà a causa dell'Alleanza, una coalizione che comanda l'intera galassia, Serenity vede come protagonista una navicella da trasporto e il suo particolare equipaggio. Malcolm (Nathan Fillion), il capitano, è ormai disilluso e incontra molte difficoltà nel seguire i propri sentimenti, con l'unica eccezione dell'affetto per i suoi compagni di avventura: Zoe (Gina Torres), primo ufficiale e moglie di Wash (Alan Tudik), il pilota; Kaylee (Jewel Staite), una ragazza molto dolce che si occupa dei motori e Jayne (Adam Baldwin) che non si separa mai da un arsenale di armi pronte all'uso. A complicare gli equilibri di questa strana famiglia arrivano Simon Tam (Sean Maher) e sua sorella River (Summer Glau): un dottore e una giovane medium ricercata dall'Alleanza a causa di un segreto, che deve rimanere tale ad ogni costo, nascosto nella sua mente.",
			  "nld: Vijfhonderd jaar in de toekomst kent de beschaving een totaal nieuwe grens en de bemanning van het ruimteschip Serenity, van de Firefly-klasse, wil een steentje bijdragen. Ze nemen elk klusje aan, of die nu legaal is of niet, om brandstof in de tank en eten op tafel te hebben. Maar de dingen worden ietwat ingewikkelder wanneer ze een passagier meenemen die wordt gezocht door het nieuwe totalitaire Alliance-regime. Zo moeten ze zelf op de vlucht slaan, alles op alles zetten om aan de Alliance-schepen en de vleesetende Reavers, die aan de rand van het heelal wonen, te ontsnappen.",
			  "pol: Kapitan Malcolm Reynolds wraz załogą statku Serenity przemierza galaktykę, przemycając ładunki lub przewożąc pasażerów. Sytuacja komplikuje się, gdy zabiera na pokład parę wyjętych spod prawa uciekinierów.",
			  "por: A tripulação de uma nave espacial viaja pelo espaço depois de os recursos da Terra terem ficado quase esgotados, obrigando-os a procurar novos recursos noutros locais, em 2517. O Capitão Malcolm \"Mal\" Reynolds orienta a sua tripulação rebelde através de tudo, tentando mantê-los vivos e a voar.",
			  "rus: Капитан Малькольм «Мэл» Рейнольдс — закаленный в боях ветеран галактической гражданской войны, сражавшийся на проигравшей, в итоге, стороне, зарабатывает на жизнь мелкими преступлениями и перевозит грузы на своем корабле «Серенити».\r\nОн возглавляет небольшую разношерстную команду, которая порядком смахивает на самую обычную семью — её члены вечно ссорятся, не соблюдают дисциплину, но никогда не предадут своего капитана и пойдут за ним хоть на край света. ",
			  "spa: Estamos alrededor del año 2500. Los humanos han colonizado multitud de planetas y satélites. La sociedad es una especie de mezcla entre el lejano oeste y la tecnología más innovadora, aunque ésta no está al alcance de todos. Los planetas centrales son los más ricos y avanzados, y los periféricos, los más atrasados. En los últimos, la gente sobrevive como puede, viviendo en planetas medio desérticos y con poca tecnología.\r\n\r\nLa Alianza (China y Estados Unidos), vencedora en una reciente guerra civil, gobierna en los sistemas, aunque su influencia en los planetas exteriores es menor. Como resultado, la cultura es una mezcla entre la china y la del salvaje oeste, y los idiomas predominantes son el chino y el inglés.\r\n\r\nNo se ha hallado vida alienígena, pero cada vez se encuentran más “reavers”, humanos enloquecidos en los confines del espacio que se dedican a asaltar las naves que encuentran, a automutilarse y otros actos terroríficos.\r\n\r\nEl Sargento Malcolm Reynolds, tras perder la guerra consigue una nave modelo Firefly, a la que llama Serenity, como el valle donde él y los demás rebeldes fueron vencidos, y forma una pequeña tripulación. Se dedican al contrabando, robos y otros delitos, aunque de vez en cuando realizan trabajos legales.",
			  "zho:         故事设定在五百年后的未来世界中，讲述由船长 Malcolm Reynolds (简称 Mal) 为首所率领的一艘没有家之太空船 Serenity于对抗联邦战争之后，在宇宙中飘流的冒险旅程。\r\n\r\n　　在太空船 Serenity 内一起生活的乘员及乘客，在穿梭於被联邦政府统治之星系中的同时，亦要在行星的边疆之间开闢新天地。不论是合法或非法，各乘员为了生存而需要接受任何的工作，而船内的每一个乘客原来也各怀鬼胎及动机，因此这班拥有著不同的过去，以及不同原因上船的傢伙们将需要在达到目的之前，各自隐藏著自已原有的身份。\r\n\r\n　　为了逃避联邦政府的追踪，Serenity 与它的乘员唯有前往许多荒无的疆界行星，亦因为这个原故，当他们面对各种挑战及困境之时，将会採用像西部牛仔般的作风去解决。FIREFLY 所谈及的是他们变化莫测的生活与关係，并於难以预料的环境下寻找著他们自已的生存意义。\r\n\r\n　　本剧虽然只播出了一季就遭砍，但经过各方的努力，终於拍摄了承接电视剧的电影版本《冲出宁静号》(Serenity)并于2005年9月30日在美国上映。",
			  "hrv: Radnja Fireflyja se odvija 500 godina u budućnosti, u novom planetarnom sustavu nakon što je čovječanstvo napustilo \"Zemlju koja je bila\". Pod vodstvom Malcoma Reynoldsa, odmetnika koji se borio protiv nove ujedninjene centralne vlasti (\"Saveza\"), posada broda Serenity, klase Firefly, na sve se načine bori za preživljavanje. Lete izmedu graničnih planeta kako bi se sakrili od Saveza."
			],
			"primary_language": "eng",
			"status": "Ended",
			"type": "series",
			"tvdb_id": "78874",
			"year": "2002",
			"slug": "firefly",
			"overviews": {
			  "ces": "Firefly je dechberoucím příběhem vesmírných psanců, kteří se pohybují na okraji společnosti, živí se převážně pašováním a jejich cílem není spasit galaxii, ale prostě přežít - \"letět dál\". Život jim komplikují jak vesmírné lodě totalitní Aliance, tak barbarští Plenitelé, sužující světy na okraji vesmíru. Hlavní síla seriálu spočívá v propracovaných charakterech postav, dobře vykonstruovaném příběhu a překvapivém rozuzlení jednotlivých epizod. Každý z hrdinů má své vlastní naděje a obavy a jejich osobnosti a vztahy se neustále vyvíjejí. ",
			  "deu": "Im 26. Jahrhundert schlägt sich die Crew des Raumfrachters \"Serenity\" vorwiegend mit Schmuggelaufträgen durch den Weltraum. Aliens wurden bisher keine entdeckt, aber Captain Malcolm \"Mal\" Reynolds, ein ehemaliger Kämpfer im verlorenen Bürgerkrieg gegen die jetzt allgegenwärtige Allianz, hat trotzdem laufend Ärger - speziell nachdem er den jungen Arzt Simon und seine Schwester River an Bord nimmt. Jene ist das geflohene \"Produkt\" von Experimenten der Allianz, und wird von dieser überall gesucht.",
			  "eng": "In the far-distant future, Captain Malcolm \"Mal\" Reynolds is a renegade former brown-coat sergeant, now turned smuggler & rogue, who is the commander of a small spacecraft, with a loyal hand-picked crew made up of the first mate, Zoe Warren; the pilot Hoban \"Wash\" Washburn; the gung-ho grunt Jayne Cobb; the engineer Kaylee Frye; the fugitives Dr. Simon Tam and his psychic sister River. Together, they travel the far reaches of space in search of food, money, and anything to live on.",
			  "fra": "Après une guerre civile qui a permis à l'Union des Planètes de dominer l'espace, le capitaine Malcolm Reynolds et son équipage s'efforcent de survivre à bord du vaisseau Serenity en effectuant diverses missions (transport de marchandises ou une mission de sauvetage...) sans trop se soucier de la légalité.",
			  "heb": "ג'וס ווידון, היוצר-תסריטאי-במאי-מפיק של \"באפי, ציידת הערפדים\" ו\"אנג'ל\", מפתיע מחדש עם סדרה שהיא הכלאה בלתי אפשרית בין מדע-בדיוני למערבון. בעתיד הקצת רחוק מסתבך הצוות של ספינת החלל \"סריניטי\" עם טרמפיסט מבוקש.",
			  "hrv": "Radnja Fireflyja se odvija 500 godina u budućnosti, u novom planetarnom sustavu nakon što je čovječanstvo napustilo \"Zemlju koja je bila\". Pod vodstvom Malcoma Reynoldsa, odmetnika koji se borio protiv nove ujedninjene centralne vlasti (\"Saveza\"), posada broda Serenity, klase Firefly, na sve se načine bori za preživljavanje. Lete izmedu graničnih planeta kako bi se sakrili od Saveza.",
			  "hun": "A Firefly egy amerikai sci-fi sorozat, amely a műfajban megszokottnál kevesebb fantasztikus, futurisztikus elemet tartalmaz, inkább hétköznapi emberi és társadalmi problémákkal teli világot mutat be, egy másik műfaj, a western témáit is felhasználva. A történet szerint a 25-26. századra a Föld készletei már rég kimerültek, az emberiség a Naprendszeren kívüli világokban telepedett le. Az űr meghódításában Kína és az USA játszott vezető szerepet, amely két szuperhatalom alkotja a Szövetséget és az Egyesítési Háborúban közös kormányzat alatt egyesítette a belső bolygókat. A külső bolygókon, sok helyen a Vadnyugatra emlékeztető, jóval rosszabb körülmények között élnek az emberek.\r\n\r\nA sorozatnak - bár csak egy évadot élt meg - hatalmas rajongótábora van.",
			  "ita": "Ambientata in un futuro privo di una reale libertà a causa dell'Alleanza, una coalizione che comanda l'intera galassia, Serenity vede come protagonista una navicella da trasporto e il suo particolare equipaggio. Malcolm (Nathan Fillion), il capitano, è ormai disilluso e incontra molte difficoltà nel seguire i propri sentimenti, con l'unica eccezione dell'affetto per i suoi compagni di avventura: Zoe (Gina Torres), primo ufficiale e moglie di Wash (Alan Tudik), il pilota; Kaylee (Jewel Staite), una ragazza molto dolce che si occupa dei motori e Jayne (Adam Baldwin) che non si separa mai da un arsenale di armi pronte all'uso. A complicare gli equilibri di questa strana famiglia arrivano Simon Tam (Sean Maher) e sua sorella River (Summer Glau): un dottore e una giovane medium ricercata dall'Alleanza a causa di un segreto, che deve rimanere tale ad ogni costo, nascosto nella sua mente.",
			  "nld": "Vijfhonderd jaar in de toekomst kent de beschaving een totaal nieuwe grens en de bemanning van het ruimteschip Serenity, van de Firefly-klasse, wil een steentje bijdragen. Ze nemen elk klusje aan, of die nu legaal is of niet, om brandstof in de tank en eten op tafel te hebben. Maar de dingen worden ietwat ingewikkelder wanneer ze een passagier meenemen die wordt gezocht door het nieuwe totalitaire Alliance-regime. Zo moeten ze zelf op de vlucht slaan, alles op alles zetten om aan de Alliance-schepen en de vleesetende Reavers, die aan de rand van het heelal wonen, te ontsnappen.",
			  "pol": "Kapitan Malcolm Reynolds wraz załogą statku Serenity przemierza galaktykę, przemycając ładunki lub przewożąc pasażerów. Sytuacja komplikuje się, gdy zabiera na pokład parę wyjętych spod prawa uciekinierów.",
			  "por": "A tripulação de uma nave espacial viaja pelo espaço depois de os recursos da Terra terem ficado quase esgotados, obrigando-os a procurar novos recursos noutros locais, em 2517. O Capitão Malcolm \"Mal\" Reynolds orienta a sua tripulação rebelde através de tudo, tentando mantê-los vivos e a voar.",
			  "rus": "Капитан Малькольм «Мэл» Рейнольдс — закаленный в боях ветеран галактической гражданской войны, сражавшийся на проигравшей, в итоге, стороне, зарабатывает на жизнь мелкими преступлениями и перевозит грузы на своем корабле «Серенити».\r\nОн возглавляет небольшую разношерстную команду, которая порядком смахивает на самую обычную семью — её члены вечно ссорятся, не соблюдают дисциплину, но никогда не предадут своего капитана и пойдут за ним хоть на край света. ",
			  "spa": "Estamos alrededor del año 2500. Los humanos han colonizado multitud de planetas y satélites. La sociedad es una especie de mezcla entre el lejano oeste y la tecnología más innovadora, aunque ésta no está al alcance de todos. Los planetas centrales son los más ricos y avanzados, y los periféricos, los más atrasados. En los últimos, la gente sobrevive como puede, viviendo en planetas medio desérticos y con poca tecnología.\r\n\r\nLa Alianza (China y Estados Unidos), vencedora en una reciente guerra civil, gobierna en los sistemas, aunque su influencia en los planetas exteriores es menor. Como resultado, la cultura es una mezcla entre la china y la del salvaje oeste, y los idiomas predominantes son el chino y el inglés.\r\n\r\nNo se ha hallado vida alienígena, pero cada vez se encuentran más “reavers”, humanos enloquecidos en los confines del espacio que se dedican a asaltar las naves que encuentran, a automutilarse y otros actos terroríficos.\r\n\r\nEl Sargento Malcolm Reynolds, tras perder la guerra consigue una nave modelo Firefly, a la que llama Serenity, como el valle donde él y los demás rebeldes fueron vencidos, y forma una pequeña tripulación. Se dedican al contrabando, robos y otros delitos, aunque de vez en cuando realizan trabajos legales.",
			  "zho": "        故事设定在五百年后的未来世界中，讲述由船长 Malcolm Reynolds (简称 Mal) 为首所率领的一艘没有家之太空船 Serenity于对抗联邦战争之后，在宇宙中飘流的冒险旅程。\r\n\r\n　　在太空船 Serenity 内一起生活的乘员及乘客，在穿梭於被联邦政府统治之星系中的同时，亦要在行星的边疆之间开闢新天地。不论是合法或非法，各乘员为了生存而需要接受任何的工作，而船内的每一个乘客原来也各怀鬼胎及动机，因此这班拥有著不同的过去，以及不同原因上船的傢伙们将需要在达到目的之前，各自隐藏著自已原有的身份。\r\n\r\n　　为了逃避联邦政府的追踪，Serenity 与它的乘员唯有前往许多荒无的疆界行星，亦因为这个原故，当他们面对各种挑战及困境之时，将会採用像西部牛仔般的作风去解决。FIREFLY 所谈及的是他们变化莫测的生活与关係，并於难以预料的环境下寻找著他们自已的生存意义。\r\n\r\n　　本剧虽然只播出了一季就遭砍，但经过各方的努力，终於拍摄了承接电视剧的电影版本《冲出宁静号》(Serenity)并于2005年9月30日在美国上映。"
			},
			"translations": {
			  "ces": "Firefly",
			  "deu": "Firefly",
			  "eng": "Firefly",
			  "fra": "Firefly",
			  "heb": "פיירפלי",
			  "hrv": "Firefly",
			  "hun": "Firefly",
			  "ita": "Firefly",
			  "kor": "파이어플라이",
			  "nld": "Firefly",
			  "pol": "Firefly",
			  "por": "Firefly",
			  "rus": "Светлячок",
			  "spa": "Firefly",
			  "swe": "Firefly",
			  "zho": "萤火虫"
			},
			"network": "FOX",
			"remote_ids": [
			  {
				"id": "tt0303461",
				"type": 0,
				"sourceName": "IMDB"
			  },
			  {
				"id": "7097",
				"type": 0,
				"sourceName": "TV.com"
			  },
			  {
				"id": "EP00524463",
				"type": 0,
				"sourceName": "Zap2It"
			  },
			  {
				"id": "1437",
				"type": 0,
				"sourceName": "TheMovieDB.com"
			  }
			],
			"thumbnail": "https://artworks.thetvdb.com/banners/posters/78874-2_t.jpg"
		  }
		]
	}`

	var fireflyResultResponse SearchResponse
	err := json.Unmarshal([]byte(fireflyResult), &fireflyResultResponse)
	if err != nil {
		t.Fatalf("Client.Search() error = %v", err)
	}

	type args struct {
		ctx           context.Context
		searchRequest SearchRequest
	}
	tests := []struct {
		name    string
		args    args
		want    SearchResponse
		wantErr bool
	}{
		{"Firefly", args{context.Background(), SearchRequest{Q: "Firefly", Limit: 1}}, fireflyResultResponse, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.Search(tt.args.ctx, tt.args.searchRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
