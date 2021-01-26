import { createPlugin } from '@backstage/core';
import CreateDeposit from './components/RecordDeposit';
import ShowDeposit from './components/DepositTable';
import SignIn from './components/SignIn';
import CreateRoomdetail from './components/RoomDetails';
import DormEmployee from './components/DormEmployee';
import CreateRepairinvoice from './components/RecordRepairinvoice';
import RepairinvoiceTablet from './components/RepairinvoiceTable';
import CreateContract from './components/RecordLease';
import ShowContract from './components/LeaseTable';
import SearchRoom from './components/SearchRoom';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
	  router.registerRoute('/DormEmployee', DormEmployee);
    router.registerRoute('/DepositTable', ShowDeposit);
    router.registerRoute('/RecordDeposit', CreateDeposit);
    router.registerRoute('/RoomDetails', CreateRoomdetail);
    router.registerRoute('/RecordRepairinvoice', CreateRepairinvoice);
    router.registerRoute('/RepairinvoiceTable', RepairinvoiceTablet );
    router.registerRoute('/RecordLease', CreateContract);
    router.registerRoute('/LeaseTable', ShowContract);
    router.registerRoute('/SearchRoom', SearchRoom);

  },
});
