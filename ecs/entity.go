package ecs

// TODO: handle corner cases

type EntityManager struct {
	queue Queue // queue of unused entity IDs
	count uint  // array of signatures where the index corresponds to the entity ID
	// total living entities
	signatures []Signature
	// TODO: consider types mapping -> entity:type
}

func (em *EntityManager) CreateEntity() Entity {
	if em.count >= uint(MaxEntities) {
		panic("Too many entities in existense")
	}

	id := em.queue.Front()
	em.queue.Remove()

	em.count++
	return id.Value.(Entity)
}

func (em *EntityManager) DestroyEntity(entity Entity) {
	// invalidate the destroyed entity's signature
	em.signatures[int(entity)] = 0
	// put the destroyed ID at the back of the queue
	em.queue.Add(entity)
	em.count--
}

func (em *EntityManager) SetSignature(entity Entity, signature Signature) {
	// put this entity's signature into the array
	em.signatures[entity] = signature
}

func (em *EntityManager) GetSignature(entity Entity) Signature {
	return em.signatures[entity]
}

func NewEntityManager() *EntityManager {
	q := NewQueue()

	// brain dead solution for now
	// TODO: consider fast queue implementation
	for i := 0; i <= MaxEntities; i++ {
		q.Add(Entity(i))
	}

	return &EntityManager{
		queue:      q,
		count:      0,
		signatures: make([]Signature, MaxEntities),
	}
}
