package fireworks

func GenerateConfiguration(fileName string) Configuration {
	configurationFilePathProvider := NewConfigurationFilePathProvider(fileName)

	configurationFileReader := NewConfigurationFileReader(configurationFilePathProvider)

	configurationJsonFileReader := NewConfigurationJsonFileReader(configurationFileReader)

	configurationJson := configurationJsonFileReader.Read()

	projectIDProvider := CreateConfigurationValueProvider("projectID", "PROJECT_ID", configurationJson)

	collectionNameProvider := CreateConfigurationValueProvider("collectionName", "COLLECTION_NAME", configurationJson)

	configurationLoader := NewApplicationConfigurationLoader(
		projectIDProvider,
		collectionNameProvider,
	)

	configuration := configurationLoader.Load()

	return configuration
}
