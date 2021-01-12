package core

const (
	STATUS_PENDING         = 0
	STATUS_ACTIVE          = 1
	STATUS_INACTIVE        = 2
	STATUS_NON_INITIALIZED = 100

	TYPE_TERMINAL     = 1
	TYPE_ROOT         = 0
	TYPE_NON_TERMINAL = 2

	NONCE_ZERO         = 0
	UNINITIALIZED_PATH = ""
	ZERO               = 0

	// TX Status constants
	TX_STATUS_PENDING    = 100
	TX_STATUS_PROCESSING = 200
	TX_STATUS_PROCESSED  = 300
	TX_STATUS_REVERTED   = 400

	BATCH_BROADCASTED = 100
	BATCH_COMMITTED   = 200

	// TODO make sure the types are correct
	TX_GENESIS           = 0
	TX_TRANSFER_TYPE     = 1
	TX_CREATE_2_TRANSFER = 3
	TX_MASS_MIGRATIONS   = 5
	TX_DEPOSIT           = 4

	CHUNK_SIZE = 3000
)
