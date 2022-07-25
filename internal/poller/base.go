package poller

func getBaseTemplate() map[string]interface{} {
	return map[string]interface{}{
		"blockhashing_blob":  "0000000000000000000000000000000000000000000000000000000000000000",
		"blocktemplate_blob": "0000000000000000000000000000000000000000000000000000000000000000",
		"difficulty":         0,
		"difficulty_top64":   0,
		"expected_reward":    0,
		"height":             0,
		"next_seed_hash":     "",
		"prev_hash":          "b44e67bdbc91e5f95c5ea49ba410b1c223ab897417b19a0f631b80cf5be43a4e",
		"reserved_offset":    0,
		"seed_hash":          "0000000000000000000000000000000000000000000000000000000000000000",
		"seed_height":        0,
		"status":             "OK",
		"untrusted":          false,
	}
}
