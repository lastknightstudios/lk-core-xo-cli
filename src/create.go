package main

func create(repository string, pipeline string, project string) {

	repo := load(repository, pipeline)

	repo.CreateRepository(project)
	//repo.CreateWebhook()
}
