//           <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Курс (1/4)</label>
//               <input
//                 type="number"
//                 placeholder="1"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.course ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.course && watch('course') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('course', {
//                   required: 'Курс обязателен',
//                   validate: (val) => {
//                     const num = Number(val)
//                     if (Number.isNaN(num)) {
//                       return 'Курс должен быть числом'
//                     }
//                     if (num < 1 || num > 4) {
//                       return 'Курс должен быть между 1 и 4'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('course')}
//               />
//               {errors.course && <p className="text-red-500 text-sm mt-1">{errors.course.message}</p>}
//             </div>

//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Семестр (1/2)</label>
//               <input
//                 type="number"
//                 placeholder="2"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.semester ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.semester && watch('semester') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('semester', {
//                   required: 'Семестр обязателен',
//                   validate: (val) => {
//                     const num = Number(val)
//                     if (Number.isNaN(num)) {
//                       return 'Семестр должен быть числом'
//                     }
//                     if (num < 1 || num > 2) {
//                       return 'Семестр должен быть 1 или 2'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('semester')}
//               />
//               {errors.semester && <p className="text-red-500 text-sm mt-1">{errors.semester.message}</p>}
//             </div>

//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Специальность</label>
//               <input
//                 type="text"
//                 placeholder="ИСП"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.specialty ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.specialty && watch('specialty') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('specialty', {
//                   required: 'Специальность обязательна',
//                   validate: (val) => {
//                     const trimmed = val.trim()
//                     if (!trimmed) {
//                       return 'Специальность не может быть пустой'
//                     }
//                     if (trimmed.length < 2) {
//                       return 'Минимум 2 символа'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('specialty')}
//               />
//               {errors.specialty && <p className="text-red-500 text-sm mt-1">{errors.specialty.message}</p>}
//             </div>

//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Выпускники (9/11)</label>
//               <input
//                 type="text"
//                 placeholder="9"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.graduation ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.graduation && watch('graduation') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('graduation', {
//                   required: 'Укажите выпуск (9 или 11)',
//                   validate: (val) =>
//                     val === '9' || val === '11' ? true : 'Выпускной класс должен быть либо 9, либо 11',
//                 })}
//                 onBlur={() => trigger('graduation')}
//               />
//               {errors.graduation && <p className="text-red-500 text-sm mt-1">{errors.graduation.message}</p>}
//             </div>

//             {/* Номер группы */}
//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Номер группы</label>
//               <input
//                 type="number"
//                 placeholder="45"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.groupNumber ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.groupNumber && watch('groupNumber') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('groupNumber', {
//                   required: 'Номер группы обязателен',
//                   pattern: {
//                     value: /^\d+$/,
//                     message: 'Номер группы должен содержать только цифры',
//                   },
//                   validate: (val) => {
//                     const num = Number(val)
//                     if (num > 999) {
//                       return 'Номер группы не может быть больше 999 (пример)'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('groupNumber')}
//               />
