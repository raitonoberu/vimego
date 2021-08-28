package vimego

type SortDirection string

const (
	AscDirection  SortDirection = "asc"
	DescDirection SortDirection = "desc"
)

type SearchFilter string

const (
	VideoFilter   SearchFilter = "clip"
	PeopleFilter  SearchFilter = "people"
	ChannelFilter SearchFilter = "channel"
	GroupFilter   SearchFilter = "group"
)

type SortOrder string

const (
	RelevanceOrder    SortOrder = "relevance"
	LatestOrder       SortOrder = "latest"
	PopularityOrder   SortOrder = "popularity"
	AlphabeticalOrder SortOrder = "alphabetical"
	DurationOrder     SortOrder = "duration"
)

type SearchCategory string

const (
	AnyCategory                    SearchCategory = ""
	TrailersCategory               SearchCategory = "trailers"
	NarrativeCategory              SearchCategory = "narrative"
	DocumentaryCategory            SearchCategory = "documentary"
	ExperimentalCategory           SearchCategory = "experimental"
	AnimationCategory              SearchCategory = "animation"
	EducationalCategory            SearchCategory = "educational"
	AdsAndCommercialsCategory      SearchCategory = "adsandcommercials"
	MusicCategory                  SearchCategory = "music"
	BrandedContentCategory         SearchCategory = "brandedcontent"
	SportsCategory                 SearchCategory = "sports"
	TravelCategory                 SearchCategory = "travel"
	CameraTechniquesCategory       SearchCategory = "cameratechniques"
	ComedyCategory                 SearchCategory = "comedy"
	EventsCategory                 SearchCategory = "events"
	FashionCategory                SearchCategory = "fashion"
	FoodCategory                   SearchCategory = "food"
	IdentsAndAnimatedLogosCategory SearchCategory = "identsandanimatedlogos"
	IndustryCategory               SearchCategory = "industry"
	IndustrionalsCategory          SearchCategory = "instructionals"
	JournalismCategory             SearchCategory = "journalism"
	PersonalCategory               SearchCategory = "personal"
	ProductCategory                SearchCategory = "product"
	TalksCategory                  SearchCategory = "talks"
	TitleAndCreditsCategory        SearchCategory = "titlesandcredits"
	VideoSchoolCategory            SearchCategory = "videoschool"
	WeedingCategory                SearchCategory = "wedding"
)
