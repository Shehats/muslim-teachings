// Only has generic state for fetch api state.
import { No_OP } from '@/store/actions'

let GenericState = {
	// action is what action to do next.
	action: No_OP,
	// observables is used to link up async observables by operation name.
	observables: {},
	// data is used to construct data to be rendered on the UI.
	data: {}
}

export default GenericState;
