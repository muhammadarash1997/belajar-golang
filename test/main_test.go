package main

import (
	"fmt"
	"testing"
)

// Cari kata cat
// Jika berawalan huruf depan maka cek
// Jika berawalan huruf belakang maka cek
// Cek kanan, Cek bawah, Cek kananbawah, Cek kiribawah

// var Reference = "dog"
// var ReferenceInverse = "god"
// var Reference = "cat"
// var ReferenceInverse = "tac"
// var Reference = "nana"
// var ReferenceInverse = "anan"
// var Reference = "huussp"
// var ReferenceInverse = "pssuuh"
var Reference = "yyqmukhmmqcsmgolykqscguhtch"
var ReferenceInverse = "hcthugcsqkylogmscqmmhkumqyy"

var Count int

func TestMain(t *testing.T) {
	// list := []string{
	// 	"aatc",
	// 	"aaaa",
	// 	"ctaa",
	// }
	// list := []string{
	// 	"gogog",
	// 	"ooooo",
	// 	"godog",
	// 	"ooooo",
	// 	"gogog",
	// }
	// list := []string{
	// 	"bananana",
	// 	"kalibrrr",
	// }
	// list := []string{
	// 	"rhhuupupsspppudusuupyspssuuuunshrphshphuusuhsssshuuhussshuusspghhsuqhuupspuusuuuuuspppssuuhssuhshn",
	// 	"jupssuspphruhsssususpsuhspuhuuhhuusspuhuuuspsssppuuhuusupssuuphuussphpuuqhuhhhhphuhppsusspspppsuuh",
	// 	"suppphpshhhhurssnuushphshpssuppsuuusyspsspuuphhhpuhsspssuuhhsuhuusspssuuuuusuuushhpjsphsushuhpuhhu",
	// 	"qspuhspsupshuhuupypusssusspsshdsuuuusppsssuupsuupsuuuhsuuhusdsjrnuuspssuupuppuuuusuusspugsphhusups",
	// 	"hshphssuuuhsshupuhuuhpupssupssuhhuuusppushpshshusuupssnuphuhushdupuupsshuhuussuussshpusspusuupsusp",
	// 	"uhuusspusshusspussuhupsujpuugussppssuuhquuuuhhusussshppuspuuspusshpspsuhsshshsssuuhssuhhsshuusspps",
	// 	"ussuuhphpuuhuhhusshhssuhsuspsphsshuupssuuhuuhhusussussuuhspuhuusuusppuhpsuuuspuupsspssppupsssshuus",
	// 	"ssuussussuupuhuuusupuhuuppsphpssuuhhuusspsuhuhsuhupphuusspuussuusususuuhhuhpphupushpususpupssuspqs",
	// 	"shsushsuhhsuuuupuupuusuhpusssuusshhshuussuhusssussssuuhspuhpsuuuhsssussuphuhusssshppssspsuhppurpds",
	// 	"pspuuuuupsuusssuhuushpssuuhsuupuupuqhuspsusypupussspppshsuhphhpupspuuhuuupphsususshuhssushhsuhsssu",
	// 	"ususuhhsusuussssuusphssupppuuspspsupsupsusphypspsspushupshhpupupshshhssussspsspuupupsspuuspssuuhsh",
	// 	"phphshsushuuhspsssuhspuussuuuuspsshsussspsyupguuuphususphuussphuusspuhsshhsusussuhsusuphuuusuupupp",
	// 	"ussusuhssuuhshpshupssuuusuduhshssppshpssshuusuuusspuususuhughupssuuhuuupuuuuusqspssyhsushsnusshssq",
	// 	"huusspususssuupuhhhphuuuhpuupssuuspuppuupssuuhhsdssuuussussuhpuupsuuphssuuuussusphpupsshpshssushyu",
	// 	"uhhhuusspupushshuspuuhshuhusssupuuuusspupphshssssuhshpuuspssuuhupsssssushshuhssussphpphhspsuhuuhhs",
	// 	"hsuphsuuppshpssuuhupuhuuhsuupuhhhsshhussphpspssuuussuusuhppspsssspspsppupspgsuhpsupuphuuupusuhpsgu",
	// 	"pssuuhsupsssuhspussssshushuussuuuspphsuuhsuphpphuhsspshuphsuspspsuuuhhpsspuhhsshuhussuhsshhhuusjuu",
	// 	"uusspphhphsuuspussshsusshsusshuushhhphqhuusspsuuhpshppusshspuuspusuuushuusuhpupsussushussuuhshshuu",
	// 	"phpusshhqujupuhpsspuuupsspspuhshupuhspuusspuhhssusspssspssuuhpshuuuushppjhpupsuhhussuuusphusshuush",
	// 	"ssspushpuushuuhphshhpsspuhpsuusppussspuhsspuuspppuussusyuhupphsuuuhsuuhupssuhpssuupsussphhssupuuhg",
	// 	"usupssuuhsuuuphhsjuhhsssspuuuupussphsuuhupupuspssuupuhuuhuhhsuusupssuuhusupspssusuhshushspspushupy",
	// 	"huphhusupspsspspuuuusuussusupsssushpughuusshsupssuhhuuusuupuhsuhuphpspsusssssussppsppspusspushuysu",
	// 	"huhsuhjssspspuhsshsppqupphudhhpshuuupuushshuuspuhuuhuhhpusssuduuuusppshhusspssshpssssuuhssupssphsu",
	// 	"upuupssuuhssnuhsusuuhssssuphpsuuuuuushspuuuuussuuuuuhuhussspuhhushsspussuuhppuyhuushpusspspuhsussp",
	// 	"uhhuuusshuupssuuhuuuhpshushpssuphhphhupuuuusppshuuhhusshssuuuupshsshuhhuhhspuspphhuppssuuhuphupuuu",
	// 	"nsrhuppsshusjuuuuuhpusuuspussphsususuhhushhsssppssuspssuuhsuspsuhpusupssuuhhssuhhhpsppsshhsusuhshu",
	// 	"huuhssshpusphuuhsphupshuusuussppssuuhusuhuusspssssshspuuuusupushuusuhspsshuhssuughusphushuuspssush",
	// 	"uuusssspspussushssushuuusssuhhrpshuuspsppuussphpspssurpshsupupusppudsunsusuuuuuuhhsussspuqpupsupsp",
	// 	"uuuspushuuhpuussshsuuuuussssuuhsspssuuhsuhuuuhuuuusshuuspushshuruupsshhpshuuuphuhuuusspuspsussnspu",
	// 	"suhuuhppssuuhppppuuuuhsuuppuphhsphssuhhpphpshspshhuuusspssunuppspusssursshpusuphuushpspushhupspspp",
	// 	"hhhuhhsuuhususssuppssuuhuhsuussupppupssuuhspuhhussuhjhssuusssuupshuuhppsususnshhhhssshpsuhhussppsu",
	// 	"puhuuhssppgupspsssspsshpususppsuqppshpushupuuuhuuphpphuusspsssspuuhusspupsuhspppuuhuusuppshushgsps",
	// 	"spupssuhupshuuhpuhuusspuppususshsssusssuushspususuussshuphuusspsqusushhpuuspsussssuuuuppsuhhuhusph",
	// 	"ssssuuspuuhuunhpuuusppsuuuspsuusssuuhuupussspsupssuuuuhuhhppsussusuuushhspssuuuushshhpssusuuuususu",
	// 	"spsruuhphususshusuhsushuspssuuhushupupuphuusssusspuusususupphphuuspshuuspupsspshussupspssuuhhuuuhu",
	// 	"ushyupsuuhussspuuhjupupupsphuussuuhusuussphuupshppssuuhguuuphuusspssuuppssuuhshhpuspuspushuuhsshps",
	// 	"uuushsuuphhspsuuuspuuhssupuphuuuhhussspsphghupuuuussphpuuhhsusuhushusshupsspuusshuphhuussuushsuuhs",
	// 	"sussushuupupuusupsshuupsshuhpphshuuuussppnusshusuuprsuudhpsssuhhpuhssussyuhuussppuuhpssuphhusppsup",
	// 	"pususshuusuuhuusspsphushpuuuhssssupspsjpsuupshppshsphhuuupuuupsuhhhshpssuuhuuphhudupupupshpssuuhsh",
	// 	"sphhupssuussshuussppusuupuuuhphjhuusspupsphpshuuhphsupsusshspssphuuuhunhuspuuuuuuhshupssphpsssuhsp",
	// 	"sushpuqssrspdssuussphsuuspsssuhuuhpsssussssuussusphhpshuhussppspuuushuhuussspssuuhspssuuhpuquhpnps",
	// 	"hhpsuusuuupppupsupsuhpssssspsshuuhhuusspusuuuhsuuusupsuhuphspupphushuuuuspssspspssuuhupsuyhupuusph",
	// 	"usuuupussspussuuhshpuuuspuuhhspuuhpssuuhhuusspuussuuuuuuspssuushpssuuhhhuuhspphpdssuuussuuhusuhhph",
	// 	"usrususspsppssuuhpsuuhuppupuuuhuupspssuuhsuuhpusupuuussuupshuusspsupssuuushdpuuhhuupshpspuhsssuuhs",
	// 	"spussphhusushhuuusussussusuhsususpupshpspuuhhshuuhhsuhsshppuuuhuhppusuphuuuspuupushssshspshpuspupp",
	// 	"spuppsuushuuuspsssusuuussshuussuusshuuupuuhsuhupushssuppphsusupusshuusspdpssuuhssuhhgpusuupssuuhyp",
	// 	"phpssupuuuuuhsupuuuduusspspupsphsshuuuuusuhuphsuusspsspspuphssuuhspsuhspppssuhsssuuhpuuhsuhshssssh",
	// 	"puhussupuhsuppuuhuphspupssphpuussuhusspspssuussjuuushssssuusususuusrspssnusspupspnussusuusuuupssyu",
	// 	"ssssuhspsshpssuuhhhhuhusshshuyuppusuhssuuuuspsusshuuhuusspsusuruhphuuuspuspuuhhshhsuphsphssuuhupjp",
	// 	"pppuuhsppssssupsppdpshsssuuuhuusshhuuupuhssupsshussuuuupssuphshuuhuussushpssuhhhsusuhhpupsphpusups",
	// 	"spssuuhuhhpusspsshuusupssuuhhhssssususuhhshpupspuuuhsqpsuhsuhuuhpuduupshsspsuuuhuupphuuusspshhhpsh",
	// 	"hpsuuuuuuusuphpsusuuuususshsuuusnhssusssuspuhspushphsspssyujhshshsshspssuusuuuhhsrspusspshssuuujup",
	// 	"uhhuuhshssushuuhpusshsuhhsssupuuupuhuuspssshuuupupsupusuhppsuusughpsssuhshsspuuusshssuspuppuhuhupd",
	// 	"pusuussuuhssuususssuhsuuhuusspsusppuphsusssusuuushsuuhsushhrsuuyuhussuuussshshpsuhuussssupupsupusp",
	// 	"puhuspuppshushupsuspupuuuuhshussusshsshspssuuhupushuusupsshquspssushhshuppudphnuuppdushussspshupps",
	// 	"ssssshhushpuppuuhuhushuuhuuspspsuupuusssuunsssuushuhhspsuushsspshpuhuusspuuuushspssuuhsuphusppssuu",
	// 	"sssuuhuusspgsshsssspssuuuygupssuuuspdspuuhussssuphhhssppuuupsnppshuussspssuuuususshuuuuhgsshusssss",
	// 	"huussppsphuuuusshhuuphpsphuupssuhhhspshsuhuppshuuhuuhhhuhsspsshhsuhupsqssshudsruhuhsuhsuhhsusupusp",
	// 	"usuhhsspssuuhspuhusshpsduuhhpsuuuuhsphspshpupuhuusuuuhsuhuusspupssushupspsshuuppshuuhsshpusuphhuhn",
	// 	"phpphpsushuphppppuspsphuuuspshuuusuhuhushuspppshphhsuhuuuushpjushuushpsssuuuuusssuuhussuuhuussupus",
	// 	"susupuususuusspsssuuushuussppssuuhususssphusssuususshuuuhpuhuhshphupsusssuguspssuuhshuspphssshsuhh",
	// 	"sussssusupsssusuuhssuppuupsspushuusussuushuhsupsuuupsuhhqusspuhssssuhuuhuhushsupssshuusupsnuussuss",
	// 	"usuhpuhuhuuuuyuhphhuspsssuhhpuphshshsusunuuuupsspsupssuuhusspuphsuhuupusuuhsusupssusquusuuhsuphusp",
	// 	"pssuuhshuhuhuuuspuhsuphshhussusspsupuuuhsuhuuussphuusspspqssppuhuupssuhhsusphshshpuppphsphpuhuuhsy",
	// 	"pphsuhsuhsuhsphuspuushuuhuuuhupspuuushpusuphhpsssupsuppssuqssuuuuhsssspsspuuuushuhsuppssuuhuhuhhph",
	// 	"usssuuuuphsuhuphpphsuuhuusspupuquurssuupushpsspsuyshusspuuupsussuuhspusphshuusspupshgssphsshuhuuus",
	// 	"ppssuuhsssspsuhsusuphsuhhhssruuuhuusssusspsssspsssuuhushssspsssuhsuuuupuushsyphushpqhupsuhshpuushp",
	// 	"uuspspsshsshpsshhuussspuuspshuusspdusspssshhpusssuuuhshuusssupuuussuuhsssssushsuuhrsuupuhuusspsphh",
	// 	"pssuuhspuupusppususssspuhphsjpsupssuuhspssuuhppuhuuusspsuspshhuuuhsphsupshsshshupsshuhhshuussssupp",
	// 	"uhjsunhsuuhhuhuususppssuuupssqphsuhuphusupsuuuhsuhusspuuuphupussuhppssuuhuuppupsuusssppssspphuuspp",
	// 	"uhuussppupssuuhssspuhsuuusuupsuupuhssdussshshspsusppuspuhpuusppusphuuuuspuususushssusuphssussssssp",
	// 	"phpssuusshuuuupsusshphpsusspsnushuussppuhssuupspsssuushussuhssssspssuuhuhspsppssuuhppusuhpuhsusspu",
	// 	"shsushussshppsspupupsssshhsusjhsquhsssuuhuuphsspsuspssuuhsusssushshsuhuphsssssdssussuuhssuspupusus",
	// 	"spsshuuussuuuhsshhsuupsusuusppssuuhsspshssuuhsuspshusuhususspsussuuhsspuspssppsuspuuhspsphuhushhsp",
	// 	"uupssuuhuhsuupuppuhpspuuhuusupungspuhuussppshuuuspuuuhpuuuuusuuuhpuppssuushuhusuhuphsupshupshuussp",
	// 	"uuusssssuusshphsussusssuhsuuuuupsssuhsspshuuuuhushhpuussuhussuuhuhsphpjuspuuhuusspuusphssusssuhssh",
	// 	"huhussushuususuppsuuushsusshuhssspuhsuspshuhhhphupsshusushhuusuuhsussuuhsuuhdphuuhshpshssupsupuuhu",
	// 	"uphuupsuusuuusssushupssuuhuspuhussuuhsupuhshshspusshususuupuhusuuusupsupusgsuuuuusupsspuhuushhhuuu",
	// 	"sspuuuhusshppshuhussuususuqsuuuuupsssshssuupussuhuushsuuphpshsssuuussuussuuhsuuhuhuphshuhpuhhupuus",
	// 	"sshpsspuhpsshuhuushpuuspsuhuuhuusspuhpussssuuuuuussjuphsshuhusppsuussssusushhpssuphphuuuusshuuuhhs",
	// 	"spuhghshjppspuuushpssssupshuusuhuspssshussuhssusupsprussuushuusshupssppuhpussussusppupusuusuqshupp",
	// 	"ussuuhsupssuuhuhhhsspshuusspupuhuhsspsuspssuushushphpuhpsuuusuusuupupppuuusupphppuussssshshuussusp",
	// 	"hsuuhsupsuuusussuuussphspupssuhuusshphuusupuhppushhhupsuuuhusuuuupushsphsppssuujhuussssjuuusspsssu",
	// 	"pushhhupssuhspuusuupsusshshhhpussuupuphrhpusuhuuppuhsuuhhuusppsussspssssshssppuusuuupuhuupspsshsps",
	// 	"ssunsuhhhhppusssupususqpssuuuussuuuushsssuuusuuhssuhpuuuupushsshspsssupssuusuhsphspupsshsushpuuphh",
	// 	"suuhspupusuuhuhhshusshuusupuuphssshhussspsuususuhssuuphshusphsusppupuhupuuhuussussphshssshssuususu",
	// 	"hpuujyusquuusssuhusshphussuuupssussuhuuuuusspsuphussuuqhsssshsppuhspssuuhhujsupshppspssupupshuuuqu",
	// 	"spsudphuussssshshpupssushuupssphpuppsuhuhuuussphuuphuhunsppsuuspsuuuhshuuhuhpshpspsussuuuusussshhs",
	// 	"hpsususuupsssshsphuusqpphhpuppshhshpsssphhhuspsuuhusuupshshususprhhhuuusphssshpssuusuuhhshuhshusus",
	// 	"yupuupqsusphpppuuhpsssusgssuhspuusuhuusspsuspssuuhshsuupphuhsshqupupsuuuhpsyssusphpususuusspppssuu",
	// 	"ssushhpspsshsjpuunuhssusjjssushuuphppsussuusuhphspussssssssuuupsyuuuuspuhupuusuupssuuhpguppsssuuus",
	// 	"hppuppuhhussuhsusspusphuuhuusspsssupupspppsssuuupupshspspssuhuspsspshrusuhhuuppususupupspsssuhshsp",
	// 	"psussduujuuusshhsnspspsuhppuuuhsshuuuspushuusppuuhspspsppsuspspshpsussuhuupshsshusquuppdspspssguus",
	// 	"ssshjpupsuuspssppssuuhhhhsspshpuhuhsspuupssuuhsuusspsppspssuphpuhpuhsuhhuussppsuhuusspsqpupppsusup",
	// 	"uusspsushshspupssuuhspussugujsssuppusushppuphuussphhhhuspphuhshpspupuuhsusuuussppphshhpssuuhnuphsh",
	// }
	list := []string{
		"msqksthghmuqcqctghshcooyhumqyh",
		"shtmccctmyshqccccscqhgqgyuhyhk",
		"qqgcmqqycscmmhmmshmcmqyykqcmmu",
		"qqksmholytqsmyqmukhmmqcsmgolyk",
		"syyykykhhukqhhockhchhumqyyquuq",
		"muyylmuuhmmqksmguqykqscguhtchy",
		"ghmulogkmsktumygmhmtuctchqgyyy",
		"ochymcgscqcummcmomuchkmhsmmsms",
		"lhsusqumshyqqshygyscmqhguucmmq",
		"ymgqmqykscyyyksqmmtyyhqmykkghk",
		"kckysqmykcykymcmshyyctqkmhcthy",
		"qyhtcysuqmqmgtguuqqsuclhsmcsll",
		"lcchqhcmhlqmsqugmogmscqmmhkmqo",
		"qcsmgouykqscmucuchqscguhtqmsyg",
		"glykmkcgyyyqmhkmmmqcsmgolykqsm",
		"usgchkmyyqmuqhkmqcsmyolyksscgs",
		"hkcmksgcsqskmohukughkyhcthugcc",
		"tkmqcctqccymkcqcmqugyyqumgcyyq",
		"cquqcsqkylqgmscsmqhkumqhhoycqm",
		"clmyoyyuoccthsgmhlysoohmclcmsm",
		"hmklyghgsoqgycmgcykyqqmyttckhh",
		"hgqoytmmhscqmqqotoyqmkmhhkhtmk",
		"cmhgcsgsthqgcktlhklmquqyuqsuhu",
		"mukccoqcsmgolykyscghhtchgsmcgm",
		"qcuqtsttmqulcuukgcthchshccyloc",
		"yymhmhhhymmgkkcqcyukstmusgchcy",
		"ymyctuuqygumcgmssshuqthmquymhy",
		"hqqtgqmgmhmmhsshqoyykhoukhqtch",
		"smhhthugcsqkylogmscqmmhkgmqyyc",
		"sysukqchkstmqcuuytqykgyclchhsh",
		"tqygqquuhcqqmyghllhyqckgohschm",
		"kmmcmusghmkkuchtoluumsqhgtqqsm",
		"ymysmmukhmmqysmcglykgscgmhtcko",
		"smgqlukqscguhlchmqygtccuscthyy",
		"okykqkkquhkcmcogssmsmyskcmmqqh",
		"umlyuhmhuhshmhcgcqmukhmqqcsmmo",
	}
	// list := []string{
	// 	"msqksthghmuqcqctghshcooyhumqyh",
	// 	"shtmccctmyshqccccscqhgqgyuhyhk",
	// 	"qqgcmqqycscmmhmmshmcmqyykqcmmu",
	// 	"qqksmholytqsmyqmukhmmqcsmgolyk",
	// 	"syyykykhhukqhhockhchhumqyyquuq",
	// 	"muyylmuuhmmqksmguqykqscguhtchy",
	// 	"ghmulogkmsktumygmhmtuctchqgyyy",
	// 	"ochymcgscqcummcmomuchkmhsmmsms",
	// 	"lhsusqumshyqqshygyscmqhguucmmq",
	// 	"ymgqmqykscyyyksqmmtyyhqmykkghk",
	// 	"kckysqmykcykymcmshyyctqkmhcthy",
	// 	"qyhtcysuqmqmgtguuqqsuclhsmcsll",
	// 	"lcchqhcmhlqmsqugmogmscqmmhkmqo",
	// 	"qcsmgouykqscmucuchqscguhtqmsyg",
	// 	"glykmkcgyyyqmhkmmmqcsmgolykqsm",
	// 	"usgchkmyyqmuqhkmqcsmyolyksscgs",
	// 	"hkcmksgcsqskmohukughkyhcthugcc",
	// 	"tkmqcctqccymkcqcmqugyyqumgcyyq",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// 	"oooooooooooooooooooooooooooooo",
	// }

	for i := 0; i < len(list); i++ {
		input := list[i]
		
		for j := 0; j < len(input); j++ {
			current := input[j]
			
			if current == Reference[0] {
				// Cek kanan
				limitRight := len(input) - len(Reference) + 1
				if j < limitRight {
					for k := 1; k < len(Reference); k++ {
						// Jika salah satu tidak sama maka jangan dilanjutkan
						if input[j+k] != Reference[k] {
							break
						}
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek kanan", Count)
						}
					}
				}

				// Cek bawah
				limitBottom := len(list) - len(Reference) + 1
				if i < limitBottom {
					for k := 1; k < len(Reference); k++ {
						// Jika salah satu tidak sama maka jangan dilanjutkan
						if list[i+k][j] != Reference[k] {
							break
						}
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek bawah", Count)
						}
					}
				}
				
				// Cek kananbawah
				if i < limitBottom && j < limitRight {
					for k := 1; k < len(Reference); k++ {
						if list[i+k][j+k] != Reference[k] {
							break
						}
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek kananbawah", Count)
						}
					}
				}
				
				// Cek kiribawah
				if i < limitBottom && j >= len(Reference)-1 {
					for k := 1; k < len(Reference); k++ {
						if list[i+k][j-k] != Reference[k] {
							break
						}
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek kiribawah", Count)
						}
					}
				}
			}

			if current == ReferenceInverse[0] {
				// Cek kanan
				limitRight := len(input) - len(ReferenceInverse) + 1
				if j < limitRight {
					for k := 1; k < len(ReferenceInverse); k++ {
						// Jika salah satu tidak sama maka jangan dilanjutkan
						if input[j+k] != ReferenceInverse[k] {
							break
						}
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek kanan inverse", Count)
						}
					}
				}

				// Cek bawah
				limitBottom := len(list) - len(Reference) + 1
				if i < limitBottom {
					for k := 1; k < len(Reference); k++ {
						// Jika salah satu tidak sama maka jangan dilanjutkan
						if list[i+k][j] != ReferenceInverse[k] {
							break
						}
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek bawah inverse", Count)
						}
					}
				}

				// Cek kananbawah
				if i < limitBottom && j < limitRight {
					for k := 1; k < len(Reference); k++ {
						if list[i+k][j+k] != ReferenceInverse[k] {
							break
						}
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek kananbawah inverse", Count)
						}
					}
				}
				// "hqqtgqmgmhmmhsshqoyykhoukhqtch",
				// "asbdbgv"
				// "asu"
				// Cek kiribawah
				fmt.Printf("Cek kiribawah, i: %d, j: %d, limitRight: %d\n", i, j, limitRight)
				if i < limitBottom && j >= len(Reference)-1 {
					fmt.Println("limitRight", limitRight)
					fmt.Println("limitRight-1", limitRight-1)
					for k := 1; k < len(Reference); k++ {
						fmt.Println("k sama dengan", k)
						fmt.Println("Masih jalan 2", i, j)
						if list[i+k][j-k] != ReferenceInverse[k] {	// error here when k 12	list[13][11-12]
							fmt.Println("Masih jalan 3", i, j)
							break
						}
						fmt.Println("Masih jalan 4", i, j)
						// Jika lolos sampai terakhir maka hitung
						if k == len(Reference)-1 {
							Count++
							fmt.Println("Cek kiribawah inverse", Count)
						}
					}
				}
			}
		}
	}

	fmt.Println(Count)
}
