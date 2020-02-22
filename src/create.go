package main

func create(repository string, pipeline string, project string) {

	repo, pipe := load(repository, pipeline)
	repo.CreateRepository(project)
	pipe.CreatePipeline(project)

}
