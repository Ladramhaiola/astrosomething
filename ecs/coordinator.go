package ecs

type Coordinator struct {
	componentManager *ComponentManger
	entiryManager    *EntityManager
	systemManager    *SystemManager
}

func (c *Coordinator) CreateEntity() Entity {
	return c.entiryManager.CreateEntity()
}

// TODO: destroy entity

func RegisterComponent[T any](c *Coordinator) {
	registerComponent[T](c.componentManager)
}

func AddComponent[T any](c *Coordinator, entity Entity, component T) {
	addComponent(c.componentManager, entity, component)

	signature := c.entiryManager.GetSignature(entity)

	signature = SetBit(signature, Signature(getComponentType[T](c.componentManager)))
	c.entiryManager.SetSignature(entity, signature)

	// notify signature changed
	c.systemManager.EntitySignatureChanged(entity, signature)
}

func RemoveComponent[T any](c *Coordinator, entity Entity) {
	removeComponent[T](c.componentManager, entity)

	signature := c.entiryManager.GetSignature(entity)

	signature = SetBit(signature, Signature(getComponentType[T](c.componentManager)))
	c.entiryManager.SetSignature(entity, signature)

	// notify signature changed
	c.systemManager.EntitySignatureChanged(entity, signature)
}

func GetComponent[T any](c *Coordinator, entity Entity) T {
	return getComponent[T](c.componentManager, entity)
}

func GetComponentType[T any](c *Coordinator) ComponentType {
	return getComponentType[T](c.componentManager)
}

func RegisterSystem(c *Coordinator, s System) {
	registerSystem(c.systemManager, s)
}

func SetSystemSignature[T System](c *Coordinator, signature Signature) {
	setSignature[T](c.systemManager, signature)
}

func NewCoordinator() *Coordinator {
	return &Coordinator{
		componentManager: NewComponentManager(),
		entiryManager:    NewEntityManager(),
		systemManager:    NewSystemManager(),
	}
}
