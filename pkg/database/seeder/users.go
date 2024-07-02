package seeder

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/helper"
	"time"
)

func SeedUsers(db *gorm.DB) error {
	var userAccountData = []*model.UserAccount{
		{Uuid: uuid.New(), Username: "john_doe", Email: "john.doe@example.com", CreatedAt: time.Now(), Password: "hashed_password_1", RefreshToken: "refresh_token_1", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "jane_smith", Email: "jane.smith@example.com", CreatedAt: time.Now(), Password: "hashed_password_2", RefreshToken: "refresh_token_2", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "michael_johnson", Email: "michael.johnson@example.com", CreatedAt: time.Now(), Password: "hashed_password_3", RefreshToken: "refresh_token_3", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "emily_wilson", Email: "emily.wilson@example.com", CreatedAt: time.Now(), Password: "hashed_password_4", RefreshToken: "refresh_token_4", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "daniel_brown", Email: "daniel.brown@example.com", CreatedAt: time.Now(), Password: "hashed_password_5", RefreshToken: "refresh_token_5", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "olivia_davis", Email: "olivia.davis@example.com", CreatedAt: time.Now(), Password: "hashed_password_6", RefreshToken: "refresh_token_6", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "william_martinez", Email: "william.martinez@example.com", CreatedAt: time.Now(), Password: "hashed_password_7", RefreshToken: "refresh_token_7", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "sophia_garcia", Email: "sophia.garcia@example.com", CreatedAt: time.Now(), Password: "hashed_password_8", RefreshToken: "refresh_token_8", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "james_rodriguez", Email: "james.rodriguez@example.com", CreatedAt: time.Now(), Password: "hashed_password_9", RefreshToken: "refresh_token_9", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "isabella_lopez", Email: "isabella.lopez@example.com", CreatedAt: time.Now(), Password: "hashed_password_10", RefreshToken: "refresh_token_10", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "liam_hernandez", Email: "liam.hernandez@example.com", CreatedAt: time.Now(), Password: "hashed_password_11", RefreshToken: "refresh_token_11", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "ava_gonzalez", Email: "ava.gonzalez@example.com", CreatedAt: time.Now(), Password: "hashed_password_12", RefreshToken: "refresh_token_12", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "noah_perez", Email: "noah.perez@example.com", CreatedAt: time.Now(), Password: "hashed_password_13", RefreshToken: "refresh_token_13", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "mia_rivera", Email: "mia.rivera@example.com", CreatedAt: time.Now(), Password: "hashed_password_14", RefreshToken: "refresh_token_14", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "logan_torres", Email: "logan.torres@example.com", CreatedAt: time.Now(), Password: "hashed_password_15", RefreshToken: "refresh_token_15", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "amelia_ramirez", Email: "amelia.ramirez@example.com", CreatedAt: time.Now(), Password: "hashed_password_16", RefreshToken: "refresh_token_16", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "ethan_flores", Email: "ethan.flores@example.com", CreatedAt: time.Now(), Password: "hashed_password_17", RefreshToken: "refresh_token_17", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "ella_sanchez", Email: "ella.sanchez@example.com", CreatedAt: time.Now(), Password: "hashed_password_18", RefreshToken: "refresh_token_18", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "alexander_nguyen", Email: "alexander.nguyen@example.com", CreatedAt: time.Now(), Password: "hashed_password_19", RefreshToken: "refresh_token_19", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "charlotte_kim", Email: "charlotte.kim@example.com", CreatedAt: time.Now(), Password: "hashed_password_20", RefreshToken: "refresh_token_20", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user1", Email: "user1@example.com", CreatedAt: time.Now(), Password: "password1", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user2", Email: "user2@example.com", CreatedAt: time.Now(), Password: "password2", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user3", Email: "user3@example.com", CreatedAt: time.Now(), Password: "password3", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user4", Email: "user4@example.com", CreatedAt: time.Now(), Password: "password4", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user5", Email: "user5@example.com", CreatedAt: time.Now(), Password: "password5", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user6", Email: "user6@example.com", CreatedAt: time.Now(), Password: "password6", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user7", Email: "user7@example.com", CreatedAt: time.Now(), Password: "password7", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user8", Email: "user8@example.com", CreatedAt: time.Now(), Password: "password8", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user9", Email: "user9@example.com", CreatedAt: time.Now(), Password: "password9", IsCompleteProfile: 1, LastLogin: nil},
		{Uuid: uuid.New(), Username: "user10", Email: "user10@example.com", CreatedAt: time.Now(), Password: "password10", IsCompleteProfile: 1, LastLogin: nil},
	}
	var userProfileData = []*model.UserProfile{
		{Uuid: uuid.New(), AccountUuid: "account_uuid_1", FirstName: "John", LastName: "Doe", Gender: "Male", DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC), Bio: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", ProfilePictureURL: "https://example.com/profile1.jpg", Longitude: 106.7985, Latitude: -6.6091},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_2", FirstName: "Jane", LastName: "Smith", Gender: "Female", DateOfBirth: time.Date(1995, time.March, 15, 0, 0, 0, 0, time.UTC), Bio: "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium.", ProfilePictureURL: "https://example.com/profile2.jpg", Longitude: 106.7973, Latitude: -6.5950},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_3", FirstName: "Michael", LastName: "Johnson", Gender: "Male", DateOfBirth: time.Date(1985, time.August, 10, 0, 0, 0, 0, time.UTC), Bio: "Phasellus vestibulum quam id pretium semper. Sed efficitur ultrices arcu, id finibus libero blandit ac.", ProfilePictureURL: "https://example.com/profile3.jpg", Longitude: 106.7961, Latitude: -6.6053},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_4", FirstName: "Emily", LastName: "Wilson", Gender: "Female", DateOfBirth: time.Date(1992, time.May, 25, 0, 0, 0, 0, time.UTC), Bio: "Nulla facilisi. Aenean posuere augue vel felis dignissim, et bibendum nisl tincidunt.", ProfilePictureURL: "https://example.com/profile4.jpg", Longitude: 106.7997, Latitude: -6.6078},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_5", FirstName: "Daniel", LastName: "Brown", Gender: "Male", DateOfBirth: time.Date(1988, time.February, 3, 0, 0, 0, 0, time.UTC), Bio: "Curabitur convallis eget velit in fermentum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae.", ProfilePictureURL: "https://example.com/profile5.jpg", Longitude: 106.7942, Latitude: -6.6021},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_6", FirstName: "Olivia", LastName: "Davis", Gender: "Female", DateOfBirth: time.Date(1991, time.October, 12, 0, 0, 0, 0, time.UTC), Bio: "Maecenas consectetur ipsum quis eros varius, non tincidunt risus pellentesque.", ProfilePictureURL: "https://example.com/profile6.jpg", Longitude: 106.8009, Latitude: -6.6015},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_7", FirstName: "William", LastName: "Martinez", Gender: "Male", DateOfBirth: time.Date(1987, time.July, 18, 0, 0, 0, 0, time.UTC), Bio: "Vestibulum mollis tempor justo, vitae pellentesque magna dictum eget.", ProfilePictureURL: "https://example.com/profile7.jpg", Longitude: 106.7980, Latitude: -6.6105},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_8", FirstName: "Sophia", LastName: "Garcia", Gender: "Female", DateOfBirth: time.Date(1993, time.January, 5, 0, 0, 0, 0, time.UTC), Bio: "Fusce tincidunt odio id ligula euismod, in fermentum elit tristique.", ProfilePictureURL: "https://example.com/profile8.jpg", Longitude: 106.7968, Latitude: -6.6039},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_9", FirstName: "James", LastName: "Rodriguez", Gender: "Male", DateOfBirth: time.Date(1984, time.December, 8, 0, 0, 0, 0, time.UTC), Bio: "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.", ProfilePictureURL: "https://example.com/profile9.jpg", Longitude: 106.7979, Latitude: -6.6083},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_10", FirstName: "Isabella", LastName: "Lopez", Gender: "Female", DateOfBirth: time.Date(1990, time.April, 20, 0, 0, 0, 0, time.UTC), Bio: "Integer at urna at elit convallis congue.", ProfilePictureURL: "https://example.com/profile10.jpg", Longitude: 106.7953, Latitude: -6.6067},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_11", FirstName: "Liam", LastName: "Hernandez", Gender: "Male", DateOfBirth: time.Date(1989, time.November, 11, 0, 0, 0, 0, time.UTC), Bio: "Vivamus vitae dolor nec nisl lobortis luctus eget id felis.", ProfilePictureURL: "https://example.com/profile11.jpg", Longitude: 106.7991, Latitude: -6.6045},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_12", FirstName: "Ava", LastName: "Gonzalez", Gender: "Female", DateOfBirth: time.Date(1994, time.February, 28, 0, 0, 0, 0, time.UTC), Bio: "Morbi lobortis purus ac nisi suscipit, a malesuada lorem congue.", ProfilePictureURL: "https://example.com/profile12.jpg", Longitude: 106.7966, Latitude: -6.6004},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_13", FirstName: "Noah", LastName: "Perez", Gender: "Male", DateOfBirth: time.Date(1986, time.September, 9, 0, 0, 0, 0, time.UTC), Bio: "Donec maximus ipsum vitae ante lacinia, nec rhoncus velit interdum.", ProfilePictureURL: "https://example.com/profile13.jpg", Longitude: 106.7987, Latitude: -6.6073},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_14", FirstName: "Mia", LastName: "Rivera", Gender: "Female", DateOfBirth: time.Date(1996, time.July, 7, 0, 0, 0, 0, time.UTC), Bio: "Suspendisse sit amet ipsum euismod, iaculis felis in, dictum risus.", ProfilePictureURL: "https://example.com/profile14.jpg", Longitude: 106.7949, Latitude: -6.6089},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_15", FirstName: "Logan", LastName: "Torres", Gender: "Male", DateOfBirth: time.Date(1983, time.March, 22, 0, 0, 0, 0, time.UTC), Bio: "Nam vel orci ac metus tincidunt ullamcorper.", ProfilePictureURL: "https://example.com/profile15.jpg", Longitude: 106.7972, Latitude: -6.6031},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_16", FirstName: "Amelia", LastName: "Ramirez", Gender: "Female", DateOfBirth: time.Date(1997, time.August, 14, 0, 0, 0, 0, time.UTC), Bio: "Phasellus auctor elit quis quam lacinia, et tincidunt velit varius.", ProfilePictureURL: "https://example.com/profile16.jpg", Longitude: 106.7995, Latitude: -6.6060},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_17", FirstName: "Ethan", LastName: "Flores", Gender: "Male", DateOfBirth: time.Date(1982, time.October, 17, 0, 0, 0, 0, time.UTC), Bio: "Quisque non ligula a justo vestibulum fermentum eget vel nunc.", ProfilePictureURL: "https://example.com/profile17.jpg", Longitude: 106.7958, Latitude: -6.6042},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_18", FirstName: "Ella", LastName: "Sanchez", Gender: "Female", DateOfBirth: time.Date(1998, time.December, 30, 0, 0, 0, 0, time.UTC), Bio: "Aliquam erat volutpat. Sed quis dolor a ante egestas consequat.", ProfilePictureURL: "https://example.com/profile18.jpg", Longitude: 106.7982, Latitude: -6.6070},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_19", FirstName: "Alexander", LastName: "Nguyen", Gender: "Male", DateOfBirth: time.Date(1981, time.January, 9, 0, 0, 0, 0, time.UTC), Bio: "Cras condimentum urna et felis lobortis, at maximus libero sollicitudin.", ProfilePictureURL: "https://example.com/profile19.jpg", Longitude: 106.7970, Latitude: -6.6025},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_20", FirstName: "Charlotte", LastName: "Kim", Gender: "Female", DateOfBirth: time.Date(1999, time.June, 5, 0, 0, 0, 0, time.UTC), Bio: "Integer vitae turpis sed ex feugiat luctus sit amet nec ligula.", ProfilePictureURL: "https://example.com/profile20.jpg", Longitude: 106.7963, Latitude: -6.6099},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "John", LastName: "Doe", Gender: "Male", DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Bio: "Love outdoor activities.", ProfilePictureURL: "http://example.com/user1.jpg", Longitude: 106.6894, Latitude: -6.5963},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Jane", LastName: "Doe", Gender: "Female", DateOfBirth: time.Date(1989, 2, 2, 0, 0, 0, 0, time.UTC), Bio: "Avid reader.", ProfilePictureURL: "http://example.com/user2.jpg", Longitude: 106.7015, Latitude: -6.5856},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Jim", LastName: "Beam", Gender: "Male", DateOfBirth: time.Date(1985, 3, 3, 0, 0, 0, 0, time.UTC), Bio: "Music lover.", ProfilePictureURL: "http://example.com/user3.jpg", Longitude: 106.7142, Latitude: -6.5748},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Jake", LastName: "Smith", Gender: "Male", DateOfBirth: time.Date(1992, 4, 4, 0, 0, 0, 0, time.UTC), Bio: "Tech enthusiast.", ProfilePictureURL: "http://example.com/user4.jpg", Longitude: 106.6831, Latitude: -6.6037},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Jill", LastName: "Johnson", Gender: "Female", DateOfBirth: time.Date(1988, 5, 5, 0, 0, 0, 0, time.UTC), Bio: "Foodie.", ProfilePictureURL: "http://example.com/user5.jpg", Longitude: 106.6948, Latitude: -6.6108},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Jerry", LastName: "Lewis", Gender: "Male", DateOfBirth: time.Date(1991, 6, 6, 0, 0, 0, 0, time.UTC), Bio: "Outdoor enthusiast.", ProfilePictureURL: "http://example.com/user6.jpg", Longitude: 106.7075, Latitude: -6.5992},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Jenny", LastName: "Lee", Gender: "Female", DateOfBirth: time.Date(1987, 7, 7, 0, 0, 0, 0, time.UTC), Bio: "Fashion lover.", ProfilePictureURL: "http://example.com/user7.jpg", Longitude: 106.6712, Latitude: -6.5885},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Jack", LastName: "Brown", Gender: "Male", DateOfBirth: time.Date(1993, 8, 8, 0, 0, 0, 0, time.UTC), Bio: "Gamer.", ProfilePictureURL: "http://example.com/user8.jpg", Longitude: 106.6839, Latitude: -6.5772},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Josie", LastName: "Williams", Gender: "Female", DateOfBirth: time.Date(1986, 9, 9, 0, 0, 0, 0, time.UTC), Bio: "Art lover.", ProfilePictureURL: "http://example.com/user9.jpg", Longitude: 106.6966, Latitude: -6.5661},
		{Uuid: uuid.New(), AccountUuid: "", FirstName: "Joe", LastName: "Davis", Gender: "Male", DateOfBirth: time.Date(1994, 10, 10, 0, 0, 0, 0, time.UTC), Bio: "Traveler.", ProfilePictureURL: "http://example.com/user10.jpg", Longitude: 106.7081, Latitude: -6.5547},
	}
	var userPreferenceData = []*model.UserPreference{
		{Uuid: uuid.New(), AccountUuid: "account_uuid_1", PreferredGender: "Female", AgeMin: 25, AgeMax: 35, InterestFor: "FITNESS", LookingFor: "LT_PARTNER", Distance: 10.5},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_2", PreferredGender: "Male", AgeMin: 30, AgeMax: 40, InterestFor: "COOKING", LookingFor: "LOOKING_FRIENDS", Distance: 15.0},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_3", PreferredGender: "Female", AgeMin: 28, AgeMax: 38, InterestFor: "OUTDOOR_ADVENTURES", LookingFor: "LOOKING_SIBLING", Distance: 12.0},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_4", PreferredGender: "Male", AgeMin: 27, AgeMax: 37, InterestFor: "MOVIES", LookingFor: "FIGURING_IT_OUT", Distance: 11.0},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_5", PreferredGender: "Female", AgeMin: 26, AgeMax: 36, InterestFor: "OUTDOOR_ADVENTURES", LookingFor: "LT_PARTNER", Distance: 13.0},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_6", PreferredGender: "Male", AgeMin: 32, AgeMax: 42, InterestFor: "ADVENTURE", LookingFor: "LOOKING_FRIENDS", Distance: 16.0},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_7", PreferredGender: "Female", AgeMin: 30, AgeMax: 40, InterestFor: "TRAVELING", LookingFor: "LOOKING_SIBLING", Distance: 14.0},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_8", PreferredGender: "Male", AgeMin: 29, AgeMax: 39, InterestFor: "FITNESS", LookingFor: "FIGURING_IT_OUT", Distance: 12.5},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_9", PreferredGender: "Female", AgeMin: 24, AgeMax: 34, InterestFor: "ART", LookingFor: "LT_PARTNER", Distance: 11.5},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_10", PreferredGender: "Male", AgeMin: 28, AgeMax: 38, InterestFor: "MUSIC_SPOTIFY", LookingFor: "LOOKING_FRIENDS", Distance: 15.5},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_11", PreferredGender: "Female", AgeMin: 26, AgeMax: 36, InterestFor: "READING", LookingFor: "LOOKING_SIBLING", Distance: 13.5},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_12", PreferredGender: "Male", AgeMin: 31, AgeMax: 41, InterestFor: "TECHNOLOGY", LookingFor: "FIGURING_IT_OUT", Distance: 17.0},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_13", PreferredGender: "Female", AgeMin: 29, AgeMax: 39, InterestFor: "FOODIE", LookingFor: "LT_PARTNER", Distance: 14.5},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_14", PreferredGender: "Male", AgeMin: 25, AgeMax: 35, InterestFor: "MOVIES", LookingFor: "LOOKING_FRIENDS", Distance: 12.8},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_15", PreferredGender: "Female", AgeMin: 27, AgeMax: 37, InterestFor: "GAMING", LookingFor: "LOOKING_SIBLING", Distance: 11.8},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_16", PreferredGender: "Male", AgeMin: 30, AgeMax: 40, InterestFor: "SPORTS", LookingFor: "FIGURING_IT_OUT", Distance: 13.8},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_17", PreferredGender: "Female", AgeMin: 28, AgeMax: 38, InterestFor: "PETS", LookingFor: "LT_PARTNER", Distance: 15.2},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_18", PreferredGender: "Male", AgeMin: 26, AgeMax: 36, InterestFor: "NATURE", LookingFor: "LOOKING_FRIENDS", Distance: 16.5},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_19", PreferredGender: "Female", AgeMin: 30, AgeMax: 40, InterestFor: "FASHION", LookingFor: "LOOKING_SIBLING", Distance: 14.7},
		{Uuid: uuid.New(), AccountUuid: "account_uuid_20", PreferredGender: "Male", AgeMin: 27, AgeMax: 37, InterestFor: "HISTORY", LookingFor: "FIGURING_IT_OUT", Distance: 12.3},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Female", AgeMin: 25, AgeMax: 35, InterestFor: "FITNESS", LookingFor: "LT_PARTNER", Distance: 4.5},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Male", AgeMin: 30, AgeMax: 40, InterestFor: "COOKING", LookingFor: "LOOKING_FRIENDS", Distance: 3.0},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Female", AgeMin: 28, AgeMax: 38, InterestFor: "OUTDOOR_ADVENTURES", LookingFor: "LOOKING_SIBLING", Distance: 4.0},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Male", AgeMin: 27, AgeMax: 37, InterestFor: "MOVIES", LookingFor: "FIGURING_IT_OUT", Distance: 4.5},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Female", AgeMin: 26, AgeMax: 36, InterestFor: "OUTDOOR_ADVENTURES", LookingFor: "LT_PARTNER", Distance: 4.0},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Male", AgeMin: 32, AgeMax: 42, InterestFor: "TRAVELING", LookingFor: "LOOKING_FRIENDS", Distance: 3.5},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Female", AgeMin: 30, AgeMax: 40, InterestFor: "TRAVELING", LookingFor: "LOOKING_SIBLING", Distance: 4.5},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Male", AgeMin: 29, AgeMax: 39, InterestFor: "FITNESS", LookingFor: "FIGURING_IT_OUT", Distance: 4.0},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Female", AgeMin: 24, AgeMax: 34, InterestFor: "ART", LookingFor: "LT_PARTNER", Distance: 3.5},
		{Uuid: uuid.New(), AccountUuid: "AccountUuid", PreferredGender: "Male", AgeMin: 28, AgeMax: 38, InterestFor: "MUSIC_SPOTIFY", LookingFor: "LOOKING_FRIENDS", Distance: 4.0},
	}

	db.Exec("DELETE FROM user_preference")
	db.Exec("DELETE FROM user_profile")
	db.Exec("DELETE FROM user_account")

	tx := db.Begin()
	for i, account := range userAccountData {
		pass, err := helper.ToHash(fmt.Sprintf("userdummy%d", i))
		if err != nil {
			tx.Rollback()
			return err
		}

		account.Email = fmt.Sprintf("userdummy%d@gmail.com", i)
		account.Password = pass

		profile := userProfileData[i]
		profile.AccountUuid = account.Uuid.String()

		preference := userPreferenceData[i]
		preference.AccountUuid = account.Uuid.String()
		if err := tx.Create(&account).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Create(&profile).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Create(&preference).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
